package domain

import (
	"context"
)

type BlockchainConfig struct {
	Difficulty   int `json:"difficulty"`
	MaxBlockSize int `json:"max_block_size"`
}

type IBlockchain interface {
	AddTransaction(ctx context.Context, tx *Transaction) error
	GetChain() []*Block
	GetMempool() []*Transaction
	GetLatestBlock() *Block
}
