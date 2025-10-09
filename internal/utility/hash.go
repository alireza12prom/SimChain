package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/alireza12prom/SimpleChain/internal/domain"
)

func hash(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func CalculateBlockHash(b *domain.Block) string {
	encodedTx, _ := json.Marshal(b.Transactions)

	record := fmt.Sprintf("%d%s%s%d%s",
		b.Index,
		b.Timestamp.Format(time.RFC3339Nano),
		b.PrevHash,
		b.Nonce,
		string(encodedTx),
	)

	return hash(record)
}

func CalculateTransactionHash(tx *domain.Transaction) string {
	record := fmt.Sprintf("%s%s%f%s",
		tx.From,
		tx.To,
		tx.Amount,
		tx.Timestamp.Format(time.RFC3339Nano),
	)

	return hash(record)
}
