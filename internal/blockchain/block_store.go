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
		if err := txn.Set(key, serialized); err != nil {
			return err
		}

		latestKey := []byte("latest_block_index")
		latestValue := []byte(fmt.Sprintf("%d", b.Index))
		return txn.Set(latestKey, latestValue)
	})
}

func (bs *BadgerStore) Close() error {
	return bs.db.Close()
}

func (bs *BadgerStore) GetLatestBlock() *domain.Block {
	var latestBlock *domain.Block

	err := bs.db.View(func(txn *badger.Txn) error {
		latestKey := []byte("latest_block_index")
		item, err := txn.Get(latestKey)
		if err != nil {
			if err == badger.ErrKeyNotFound {
				return nil
			}
			return err
		}

		var latestIndex int
		err = item.Value(func(val []byte) error {
			_, err := fmt.Sscanf(string(val), "%d", &latestIndex)
			return err
		})
		if err != nil {
			return err
		}

		blockKey := []byte(fmt.Sprintf("block_%d", latestIndex))
		blockItem, err := txn.Get(blockKey)
		if err != nil {
			return err
		}

		return blockItem.Value(func(val []byte) error {
			block, err := utils.DeserializeBlock(val)
			if err != nil {
				return err
			}
			latestBlock = block
			return nil
		})
	})
	if err != nil {
		return nil
	}

	return latestBlock
}

func (bs *BadgerStore) GetChain() ([]*domain.Block, error) {
	var chain []*domain.Block
	blocks := make(map[int]*domain.Block)
	var maxIndex int = -1

	err := bs.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			key := item.Key()

			if len(key) >= 6 && string(key[:6]) == "block_" {
				var index int
				if _, err := fmt.Sscanf(string(key), "block_%d", &index); err == nil {
					if index > maxIndex {
						maxIndex = index
					}

					err := item.Value(func(val []byte) error {
						block, err := utils.DeserializeBlock(val)
						if err != nil {
							return err
						}
						blocks[index] = block
						return nil
					})
					if err != nil {
						return err
					}
				}
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	// Build chain in order from index 0 to maxIndex
	for i := 0; i <= maxIndex; i++ {
		if block, exists := blocks[i]; exists {
			chain = append(chain, block)
		}
	}

	return chain, nil
}

var _ domain.IBlockStore = (*BadgerStore)(nil)
