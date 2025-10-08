package domain

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/alireza12prom/SimpleChain/internal/utility"
)

type Transaction struct {
	Hash      string  `json:"id"`
	From      string  `json:"from"`
	To        string  `json:"to"`
	Amount    float64 `json:"amount"`
	Timestamp int64   `json:"timestamp"`
}

func NewTransaction(from, to string, amount float64) *Transaction {
	tx := &Transaction{
		From:      from,
		To:        to,
		Amount:    amount,
		Timestamp: time.Now().Unix(),
	}

	tx.Hash = tx.GetHash()
	return tx
}

func (block *Transaction) GetHash() string {
	data, _ := json.Marshal(block)
	return utility.GetHash(string(data))
}

func (tx *Transaction) Validate() error {
	if tx.From == "" || tx.To == "" {
		return fmt.Errorf("from or to address is empty")
	}
	if tx.Amount <= 0 {
		return fmt.Errorf("amount must be positive")
	}
	if tx.Hash != tx.GetHash() {
		return fmt.Errorf("invalid transaction hash")
	}
	return nil
}

func (tx *Transaction) ToJSON() ([]byte, error) {
	return json.Marshal(tx)
}
