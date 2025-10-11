package blockchain

import (
	"context"

	"github.com/alireza12prom/SimpleChain/internal/domain"
)

type Blockchain struct {
	config  domain.BlockchainConfig
	memPool *TransactionPool
}

func NewBlockchain(config domain.BlockchainConfig) *Blockchain {
	return &Blockchain{config: config, memPool: NewTransactionPool()}
}

func (bc *Blockchain) AddBlock(ctx *context.Context, block *domain.Block) error {
	return nil
}

func (bc *Blockchain) AddTransaction(ctx *context.Context, tx *domain.Transaction) error {
	return bc.memPool.AddTransaction(tx)
}

func (bc *Blockchain) ValidateBlock(ctx *context.Context, b *domain.Block) error {
	return nil
}

func (bc *Blockchain) GetLatestBlock() *domain.Block {
	return nil
}

func (bc *Blockchain) GetMempool() []*domain.Transaction {
	return nil
}

func (bc *Blockchain) GetChain() []*domain.Block {
	return nil
}

func (bc *Blockchain) ValidateChain() error {
	return nil
}

var _ domain.IBlockchain = (*Blockchain)(nil)
