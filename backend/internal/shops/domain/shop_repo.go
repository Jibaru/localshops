package domain

import (
	"context"
	"errors"
)

var (
	ErrShopRepoCanNotSave = errors.New("can not save shop")
	ErrShopRepoCanNotGet  = errors.New("can not get")
)

type ShopRepo interface {
	NextID() string
	Save(ctx context.Context, shop *Shop) error
	Get(ctx context.Context) ([]*Shop, error)
}
