package core

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"

	"github.com/dgraph-io/badger/v4"
)

type Blockchain struct {
	Chain      []Block `json:"chain"`
	Difficulty int     `json:"difficulty"`
	DB         *badger.DB
}

func NewBlockchain(db *badger.DB) *Blockchain {
	return &Blockchain{
		Chain:      []Block{},
		Difficulty: 4,
		DB:         db,
	}
}

func (bc *Blockchain) Init() {
	tx := Transaction{
		Hash:      "genesis",
		From:      "system",
		To:        "genesis",
		Amount:    0,
		Timestamp: time.Now().Unix(),
	}

	block := Block{
		Index:        0,
		Timestamp:    time.Now(),
		Transactions: []Transaction{tx},
		PrevHash:     "0",
		Nonce:        0,
	}

	bc.Chain = append(bc.Chain, block)
	block.Hash = bc.calculateHash(block)

	bc.saveBlock(genesisBlock)
}

func (bc *Blockchain) calculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp.String() + block.PrevHash + strconv.Itoa(block.Nonce)

	for _, tx := range block.Transactions {
		record += tx.Hash + tx.From + tx.To + fmt.Sprintf("%.2f", tx.Amount)
	}

	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
}
