package core

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/dgraph-io/badger/v4"
)

type Blockchain struct {
	Difficulty int `json:"difficulty"`
	DB         *badger.DB
	LastBlock  *Block
}

func NewBlockchain(db *badger.DB) *Blockchain {
	return &Blockchain{
		Difficulty: 4,
		DB:         db,
	}
}

func (bc *Blockchain) Init() error {
	err := bc.DB.Update(func(txn *badger.Txn) error {
		_, err := txn.Get([]byte("Block:0"))
		if err != nil {
			if err.Error() == "ErrKeyNotFound" {
				transaction := NewTransaction("__system", "__system", 0)
				bc.LastBlock = NewBlock(1, "0")
				bc.LastBlock.AddTransaction(transaction)
				bc.save(txn)
			} else {
				return err
			}
		}
		return err
	})
	if err != nil {
		return err
	}

	return nil
}

func (bc *Blockchain) calculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp.String() + block.PrevHash + strconv.Itoa(block.Nonce)

	for _, tx := range block.Transactions {
		record += tx.Hash + tx.From + tx.To + fmt.Sprintf("%.2f", tx.Amount)
	}

	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
}

func (bc *Blockchain) save(txn *badger.Txn) {
	txn.Set([]byte(fmt.Sprintf("block:%s", bc.LastBlock.Hash)), []byte(bc.LastBlock.ToJSON()))
}
