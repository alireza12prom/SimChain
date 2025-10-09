package domain

import (
	"context"
)

type IBlockStore interface {
	SaveBlock(ctx context.Context, b *Block) error
	LoadChain(ctx context.Context) ([]*Block, error)
	Close() error
}
