package blockchain

import (
	"context"

	"github.com/alireza12prom/SimpleChain/internal/domain"
)

type Blockchain struct {
	config     domain.BlockchainConfig
	memPool    *TransactionPool
	blockStore domain.IBlockStore
}

func NewBlockchain(config domain.BlockchainConfig, blockStore domain.IBlockStore) *Blockchain {
	return &Blockchain{
		config:     config,
		memPool:    NewTransactionPool(),
		blockStore: blockStore,
	}
}

func (bc *Blockchain) AddTransaction(ctx context.Context, tx *domain.Transaction) error {
	return bc.memPool.AddTransaction(tx)
}

func (bc *Blockchain) GetLatestBlock() *domain.Block {
	return bc.blockStore.GetLatestBlock()
}

func (bc *Blockchain) GetMempool() []*domain.Transaction {
	return bc.memPool.GetTransactions()
}

func (bc *Blockchain) GetChain() []*domain.Block {
	chain, err := bc.blockStore.GetChain()
	if err != nil {
		return nil
	}
	return chain
}

var _ domain.IBlockchain = (*Blockchain)(nil)
