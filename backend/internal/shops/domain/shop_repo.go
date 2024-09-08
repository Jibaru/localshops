package domain

import (
	"context"
	"errors"
)

var (
	ErrShopRepoCanNotSave = errors.New("can not save shop")
)

type ShopRepo interface {
	NextID() string
	Save(ctx context.Context, shop *Shop) error
}
