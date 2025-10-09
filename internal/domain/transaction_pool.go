package domain

type ITransactionPool interface {
	AddTransaction(*Transaction) error
	GetTransactions() []*Transaction
	RemoveTransaction(*Transaction) error
}
