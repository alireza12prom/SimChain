package domain

type ITransactionPool interface {
	AddTransaction(*Transaction) error
	GetTransactions(n int) []*Transaction
	GetPool() []*Transaction
	RemoveTransaction(...*Transaction) error
}
