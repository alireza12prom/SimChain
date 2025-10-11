package domain

import (
	"context"
)

type IBlockStore interface {
	SaveBlock(ctx context.Context, b *Block) error
	Close() error
	GetLatestBlock() *Block
	GetChain() ([]*Block, error)
}
