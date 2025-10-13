package blockchain

import (
	"slices"
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

func (tp *TransactionPool) GetTransactions(n int) []*domain.Transaction {
	tp.mu.RLock()
	defer tp.mu.RUnlock()

	txs := make([]*domain.Transaction, 0, n)
	for i := 0; i < n; i++ {
		if i >= len(tp.pool) {
			break
		}
		txs = append(txs, tp.pool[i])
	}

	return txs
}

func (tp *TransactionPool) GetPool() []*domain.Transaction {
	tp.mu.RLock()
	defer tp.mu.RUnlock()
	return append([]*domain.Transaction{}, tp.pool...)
}

func (tp *TransactionPool) RemoveTransaction(txs ...*domain.Transaction) error {
	tp.mu.Lock()
	defer tp.mu.Unlock()
	tp.pool = slices.DeleteFunc(tp.pool, func(t *domain.Transaction) bool {
		return slices.Contains(txs, t)
	})
	return nil
}

var _ domain.ITransactionPool = (*TransactionPool)(nil)
