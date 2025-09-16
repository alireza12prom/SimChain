package core

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Transaction struct {
	Hash      string  `json:"id"`
	From      string  `json:"from"`
	To        string  `json:"to"`
	Amount    float64 `json:"amount"`
	Fee       float64 `json:"fee"`
	Nonce     uint64  `json:"nonce"`
	Signature string  `json:"signature"`
	Version   int     `json:"version"`
	Timestamp int64   `json:"timestamp"`
}

func NewTransaction(from, to string, amount, fee float64, nonce uint64, data, chainID string) *Transaction {
	tx := &Transaction{
		From:      from,
		To:        to,
		Amount:    amount,
		Fee:       fee,
		Nonce:     nonce,
		Timestamp: time.Now().Unix(),
		Version:   1,
	}

	tx.Hash = tx.calculateHash()
	return tx
}

// calculateHash generates the transaction's hash based on its contents
func (tx *Transaction) calculateHash() string {
	data := fmt.Sprintf(
		"%s%s%.8f%.8f%d%d",
		tx.From,
		tx.To,
		tx.Amount,
		tx.Fee,
		tx.Nonce,
		tx.Timestamp,
	)

	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func (*Transaction) Sign() {}

func (*Transaction) VerifySignature() {}

func (*Transaction) ToJson() {}

func (*Transaction) FromJson() {}
