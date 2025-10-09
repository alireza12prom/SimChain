package domain

import (
	"context"
)

type BlockchainConfig struct {
	Difficulty   int `json:"difficulty"`
	MaxBlockSize int `json:"max_block_size"`
}

type IBlockchain interface {
	AddBlock(ctx *context.Context, block *Block) error
	AddTransaction(ctx *context.Context, tx *Transaction) error
	GetChain() []*Block
	ValidateChain() error
	ValidateBlock(ctx *context.Context, b *Block) error
	GetMempool() []*Transaction
	GetLatestBlock() *Block
}
