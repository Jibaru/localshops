package domain

import (
	"context"
	"errors"
)

var (
	ErrProductRepoCanNotSave = errors.New("can not save product")
)

type ProductRepo interface {
	NextID() string
	Save(ctx context.Context, product *Product) error
}
