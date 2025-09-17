package core

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"
)

type Transaction struct {
	Hash      string  `json:"id"`
	From      string  `json:"from"`
	Symbol    string  `json:"symbol"`
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

	tx.Hash = tx.calculateHash()
	return tx
}

func (tx *Transaction) calculateHash() string {
	data := fmt.Sprintf(
		"%s%s%.8f%.8f%d%d",
		tx.From,
		tx.To,
		tx.Amount,
		tx.Timestamp,
	)

	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func (tx *Transaction) Validate() error {
	if tx.From == "" || tx.To == "" {
		return fmt.Errorf("from or to address is empty")
	}
	if tx.Amount <= 0 {
		return fmt.Errorf("amount must be positive")
	}
	if tx.Symbol == "" {
		return fmt.Errorf("symbol cannot be empty")
	}
	if tx.Hash != tx.calculateHash() {
		return fmt.Errorf("invalid transaction hash")
	}
	return nil
}

func (tx *Transaction) ToJSON() ([]byte, error) {
	return json.Marshal(tx)
}
