package blockchain

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/alireza12prom/SimpleChain/internal/domain"
	utils "github.com/alireza12prom/SimpleChain/internal/utility"
)

type Blockchain struct {
	config     domain.BlockchainConfig
	memPool    *TransactionPool
	blockStore domain.IBlockStore
}

func NewBlockchain(config domain.BlockchainConfig, blockStore domain.IBlockStore) *Blockchain {
	bc := &Blockchain{config: config, memPool: NewTransactionPool(), blockStore: blockStore}

	if blockStore.GetLatestBlock() == nil {
		genesis := &domain.Block{
			Index:        0,
			Timestamp:    time.Now(),
			Transactions: nil,
			PrevHash:     "",
			Nonce:        0,
		}
		genesis.Hash = utils.CalculateBlockHash(genesis)
		_ = blockStore.SaveBlock(context.Background(), genesis)
	}

	return bc
}

func (bc *Blockchain) CreateBlock() (*domain.Block, error) {
	txs := bc.memPool.GetTransactions(bc.config.MaxBlockSize)

	if len(txs) == 0 {
		return nil, errors.New("no transactions to create block")
	}

	block := &domain.Block{
		Index:        bc.GetLatestBlock().Index + 1,
		Transactions: txs,
		Timestamp:    time.Now(),
		PrevHash:     bc.GetLatestBlock().Hash,
		Nonce:        0,
	}

	prefix := strings.Repeat("0", bc.config.Difficulty)
	for {
		block.Hash = utils.CalculateBlockHash(block)
		if strings.HasPrefix(block.Hash, prefix) {
			break
		}
		block.IncreaseNonce()
	}

	bc.blockStore.SaveBlock(context.Background(), block)
	bc.memPool.RemoveTransaction(txs...)

	return block, nil
}

func (bc *Blockchain) AddTransaction(ctx context.Context, from string, to string, amount float64) error {
	tx := &domain.Transaction{
		From:      from,
		To:        to,
		Amount:    amount,
		Timestamp: time.Now(),
	}
	tx.Hash = utils.CalculateTransactionHash(tx)
	return bc.memPool.AddTransaction(tx)
}

func (bc *Blockchain) GetLatestBlock() *domain.Block {
	return bc.blockStore.GetLatestBlock()
}

func (bc *Blockchain) GetMempool() []*domain.Transaction {
	return bc.memPool.GetPool()
}

func (bc *Blockchain) GetChain() []*domain.Block {
	chain, err := bc.blockStore.GetChain()
	if err != nil {
		return nil
	}
	return chain
}

var _ domain.IBlockchain = (*Blockchain)(nil)
