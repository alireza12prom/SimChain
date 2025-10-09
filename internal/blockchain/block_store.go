package blockchain

import (
	"context"
	"fmt"

	"github.com/alireza12prom/SimpleChain/internal/domain"
	utils "github.com/alireza12prom/SimpleChain/internal/utility"
	"github.com/dgraph-io/badger/v4"
)

type BadgerStore struct {
	db *badger.DB
}

func NewBadgerStore(dataDir string) domain.IBlockStore {
	opts := badger.DefaultOptions(dataDir)
	db, err := badger.Open(opts)
	if err != nil {
		panic(err)
	}
	return &BadgerStore{db: db}
}

func (bs *BadgerStore) SaveBlock(ctx context.Context, b *domain.Block) error {
	serialized, err := utils.SerializeBlock(b)
	if err != nil {
		return err
	}
	key := []byte(fmt.Sprintf("block_%d", b.Index))

	return bs.db.Update(func(txn *badger.Txn) error {
		return txn.Set(key, serialized)
	})
}

func (bs *BadgerStore) LoadChain(ctx context.Context) ([]*domain.Block, error) {
	var chain []*domain.Block
	return chain, nil
}

func (bs *BadgerStore) Close() error {
	return bs.db.Close()
}

var _ domain.IBlockStore = (*BadgerStore)(nil)
