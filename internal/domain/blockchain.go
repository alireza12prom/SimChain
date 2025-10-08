package domain

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/dgraph-io/badger/v4"
)

type Blockchain struct {
	Difficulty   int `json:"difficulty"`
	DB           *badger.DB
	LastBlock    *Block
	Mempool      []*Transaction
	MaxBlockSize int
}

func (bc *Blockchain) save(txn *badger.Txn, block *Block) error {
	key := fmt.Sprintf("block:%d", block.Index)
	value := []byte(block.ToJSON())

	if err := txn.Set([]byte(key), value); err != nil {
		return err
	}

	lastIndexKey := []byte("last_block_index")
	lastIndexValue := []byte(strconv.Itoa(block.Index))
	if err := txn.Set(lastIndexKey, lastIndexValue); err != nil {
		return err
	}

	bc.LastBlock = block
	return nil
}

func (bc *Blockchain) removeMinedTransactions(block *Block) {
	minedTrxHash := make(map[string]bool)
	for _, tx := range block.Transactions {
		minedTrxHash[tx.Hash] = true
	}

	newMempool := []*Transaction{}
	for _, tx := range bc.Mempool {
		if !minedTrxHash[tx.Hash] {
			newMempool = append(newMempool, tx)
		}
	}

	bc.Mempool = newMempool
}

func (bc *Blockchain) loadLastBlock(txn *badger.Txn) error {
	lastIndexKey := []byte("last_block_index")
	lastBlockIndex, err := txn.Get(lastIndexKey)
	if err != nil {
		return err
	}

	var lastBlockKey []byte
	err = lastBlockIndex.Value(func(val []byte) error {
		lastBlockKey = []byte(fmt.Sprintf("block:%d", string(val)))
		return err
	})
	if err != nil {
		return err
	}

	lastBlockValue, err := txn.Get(lastBlockKey)
	if err != nil {
		return err
	}

	err = lastBlockValue.Value(func(val []byte) error {
		lastBlock, err := BlockFromJSON(val)
		if err != nil {
			return nil
		}
		bc.LastBlock = lastBlock
		return nil
	})

	return err
}

func NewBlockchain(db *badger.DB) *Blockchain {
	return &Blockchain{
		Difficulty:   4,
		DB:           db,
		Mempool:      []*Transaction{},
		MaxBlockSize: 10,
	}
}

func (bc *Blockchain) Init() error {
	err := bc.DB.View(func(txn *badger.Txn) error {
		_, err := txn.Get([]byte("block:0"))
		if err != nil {
			if err == badger.ErrKeyNotFound {
				transaction := NewTransaction("__system", "__system", 0)
				bc.LastBlock = NewBlock(1, "0")
				bc.LastBlock.AddTransaction(transaction)
				bc.save(txn, bc.LastBlock)

				return nil
			}
			return err
		} else {
			err = bc.loadLastBlock(txn)
			if err != nil {
				return err
			}
		}
		return err
	})
	return err
}

func (bc *Blockchain) AddTransaction(trx *Transaction) {
	bc.Mempool = append(bc.Mempool, trx)
}

func (bc *Blockchain) MineBlock() (*Block, error) {
	if len(bc.Mempool) == 0 {
		return nil, errors.New("no transaction in pool to min")
	}

	block := NewBlock(bc.LastBlock.Index+1, bc.LastBlock.Hash)

	for _, tx := range bc.Mempool {
		block.AddTransaction(tx)
		if block.Size() >= bc.MaxBlockSize {
			break
		}
	}

	fmt.Printf("🔨 Mining block with %d transactions\n", block.Size())

	target := ""
	for i := 0; i < bc.Difficulty; i++ {
		target += "0"
	}

	for {
		if block.Hash[:bc.Difficulty] == target {
			fmt.Printf("🚚 Block mined: %s\n", block.Hash)
			break
		}
		block.IncreaseNone()
	}

	// Save block
	err := bc.DB.Update(func(txn *badger.Txn) error {
		return bc.save(txn, block)
	})
	if err != nil {
		return nil, err
	}

	bc.removeMinedTransactions(block)

	return block, nil
}

func (bc *Blockchain) GetPendingTransaction() []*Transaction {
	return bc.Mempool
}

func (bc *Blockchain) GetBlocks() []*Block {
	blocks := []*Block{}

	err := bc.DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("last_block_index"))
		if err != nil {
			return err
		}

		var lastIndex int
		err = item.Value(func(val []byte) error {
			lastIndex, err = strconv.Atoi(string(val))
			return err
		})
		if err != nil {
			return err
		}

		// Load all blocks from 0 to lastIndex
		for i := 0; i <= lastIndex; i++ {
			key := fmt.Sprintf("block:%d", i)
			item, err := txn.Get([]byte(key))
			if err != nil {
				// Skip missing blocks (shouldn't happen in valid chain)
				continue
			}

			err = item.Value(func(val []byte) error {
				block, err := BlockFromJSON(val)
				if err != nil {
					return err
				}

				blocks = append(blocks, block)
				return nil
			})
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		fmt.Printf("Error loading blocks: %v\n", err)
		return []*Block{}
	}

	return blocks
}
