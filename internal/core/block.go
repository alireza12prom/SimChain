package core

import (
	"encoding/json"
	"time"

	"github.com/alireza12prom/SimpleChain/internal/utility"
)

type Block struct {
	Index        int            `json:"index"`
	Timestamp    time.Time      `json:"timestamp"`
	Transactions []*Transaction `json:"transactions"`
	PrevHash     string         `json:"prev_hash"`
	Hash         string         `json:"hash"`
	Nonce        int            `json:"nonce"`
}

func NewBlock(index int, prev_hash string) *Block {
	return &Block{
		Index:        index,
		Timestamp:    time.Now(),
		Transactions: []*Transaction{},
		PrevHash:     prev_hash,
		Nonce:        0,
	}
}

func (block *Block) GetHash() string {
	data, _ := json.Marshal(block)
	return utility.GetHash(string(data))
}

func (b *Block) AddTransaction(trx *Transaction) {
	b.Transactions = append(b.Transactions, trx)
}

func (b *Block) ToJSON() string {
	value, _ := json.Marshal(b)
	return string(value[:])
}
