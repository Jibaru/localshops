package domain

import (
	"context"
	"errors"
)

var (
	ErrShopRepoCanNotSave    = errors.New("can not save shop")
	ErrShopRepoCanNotGet     = errors.New("can not get")
	ErrShopRepoCanNotGetByID = errors.New("can not get shop by id")
)

type ShopRepo interface {
	NextID() string
	Save(ctx context.Context, shop *Shop) error
	Get(ctx context.Context) ([]*Shop, error)
	GetByID(ctx context.Context, id string) (*Shop, error)
}
