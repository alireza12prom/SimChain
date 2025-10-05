package core

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

	return txn.Set(lastIndexKey, lastIndexValue)
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
		_, err := txn.Get([]byte("Block:0"))
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
		}
		return err
	})
	if err != nil {
		return err
	}

	return nil
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
