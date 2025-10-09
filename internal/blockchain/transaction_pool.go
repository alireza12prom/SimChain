package blockchain

import (
	"sync"

	"github.com/alireza12prom/SimpleChain/internal/domain"
)

type TransactionPool struct {
	pool []*domain.Transaction
	mu   sync.RWMutex
}

func NewTransactionPool() *TransactionPool {
	return &TransactionPool{pool: []*domain.Transaction{}}
}

func (tp *TransactionPool) AddTransaction(tx *domain.Transaction) error {
	tp.mu.Lock()
	defer tp.mu.Unlock()
	tp.pool = append(tp.pool, tx)
	return nil
}

func (tp *TransactionPool) GetTransactions() []*domain.Transaction {
	tp.mu.RLock()
	defer tp.mu.RUnlock()
	return append([]*domain.Transaction{}, tp.pool...)
}

func (tp *TransactionPool) RemoveTransaction(tx *domain.Transaction) error {
	tp.mu.Lock()
	defer tp.mu.Unlock()
	for i, t := range tp.pool {
		if t == tx {
			tp.pool = append(tp.pool[:i], tp.pool[i+1:]...)
			return nil
		}
	}
	return nil
}

var _ domain.ITransactionPool = (*TransactionPool)(nil)
