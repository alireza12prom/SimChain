package domain

import (
	"time"
)

type Block struct {
	Index        int            `json:"index"`
	Timestamp    time.Time      `json:"timestamp"`
	Transactions []*Transaction `json:"transactions"`
	PrevHash     string         `json:"prev_hash"`
	Hash         string         `json:"hash"`
	Nonce        int            `json:"nonce"`
}

type IBlock interface {
	IncreaseNonce()
	AddTransaction(trx *Transaction)
}

func (b *Block) IncreaseNonce() {
	b.Nonce += 1
}

func (b *Block) AddTransaction(trx *Transaction) {
	b.Transactions = append(b.Transactions, trx)
}

var _ IBlock = (*Block)(nil)
