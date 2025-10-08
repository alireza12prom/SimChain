package domain

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
	block := &Block{
		Index:        index,
		Timestamp:    time.Now(),
		Transactions: []*Transaction{},
		PrevHash:     prev_hash,
		Nonce:        0,
	}

	block.Hash = block.GetHash()
	return block
}

func (b *Block) GetHash() string {
	data, _ := json.Marshal(b)
	return utility.GetHash(string(data))
}

func (b *Block) AddTransaction(trx *Transaction) {
	b.Transactions = append(b.Transactions, trx)
	b.Hash = b.GetHash()
}

func (b *Block) Size() int {
	return len(b.Transactions)
}

func (b *Block) IncreaseNone() {
	b.Nonce += 1
	b.Hash = b.GetHash()
}

func (b *Block) ToJSON() string {
	value, _ := json.Marshal(b)
	return string(value[:])
}

func BlockFromJSON(data []byte) (*Block, error) {
	block := &Block{}

	if err := json.Unmarshal(data, &block); err != nil {
		return nil, err
	}

	return block, nil
}
