package domain

import (
	"context"
	"errors"
)

var (
	ErrRemoteShopServExistsByID = errors.New("shop exists by id failed")
)

type RemoteShopServ interface {
	ExistsByID(ctx context.Context, id string) (bool, error)
}
