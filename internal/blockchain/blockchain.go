package blockchain

import (
	"context"

	"github.com/alireza12prom/SimpleChain/internal/domain"
)

type Blockchain struct {
	config domain.BlockchainConfig
}

func (bc *Blockchain) AddBlock(ctx *context.Context, block *domain.Block) error {
}

func (bc *Blockchain) ValidateBlock(ctx *context.Context, b *domain.Block) error {
}

func (bc *Blockchain) AddTransaction(ctx *context.Context, tx *domain.Transaction) error {
}

func (bc *Blockchain) GetLatestBlock() *domain.Block {
}

func (bc *Blockchain) GetMempool() []*domain.Transaction {
}

func (bc *Blockchain) GetChain() []*domain.Block {
}

func (bc *Blockchain) ValidateChain() error {
}

var _ domain.IBlockchain = (*Blockchain)(nil)
