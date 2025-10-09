package utils

import (
	"bytes"
	"encoding/gob"

	"github.com/alireza12prom/SimpleChain/internal/domain"
)

func SerializeBlock(b *domain.Block) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(b); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func DeserializeBlock(data []byte) (*domain.Block, error) {
	var b domain.Block
	dec := gob.NewDecoder(bytes.NewReader(data))
	if err := dec.Decode(&b); err != nil {
		return nil, err
	}
	return &b, nil
}

func SerializeTransaction(tx *domain.Transaction) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(tx); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func DeserializeTransaction(data []byte) (*domain.Transaction, error) {
	var tx domain.Transaction
	dec := gob.NewDecoder(bytes.NewReader(data))
	if err := dec.Decode(&tx); err != nil {
		return nil, err
	}
	return &tx, nil
}

func SerializeChain(chain []*domain.Block) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(chain); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func DeserializeChain(data []byte) ([]*domain.Block, error) {
	var chain []*domain.Block
	dec := gob.NewDecoder(bytes.NewReader(data))
	if err := dec.Decode(&chain); err != nil {
		return nil, err
	}
	return chain, nil
}
