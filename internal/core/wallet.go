package core

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Wallet struct {
	Address    string            `json:"address"`
	PrivateKey *ecdsa.PrivateKey `json:"-"`
	PublicKey  *ecdsa.PublicKey  `json:"public_key"`
	Balance    float64           `json:"balance"`
}

func NewWallet() (*Wallet, error) {
	wallet := &Wallet{
		Balance: 0.0,
	}

	private, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("failed to generate private key: %v", err)
	}
	wallet.PrivateKey = private
	wallet.PublicKey = &private.PublicKey

	publicKeyBytes := append(wallet.PublicKey.X.Bytes(), wallet.PrivateKey.Y.Bytes()...)
	hash := sha256.Sum256(publicKeyBytes)
	wallet.Address = hex.EncodeToString(hash[:])

	return wallet, nil
}

func (w *Wallet) GetBalance() {}

func (w *Wallet) UpdateBalance() {}
