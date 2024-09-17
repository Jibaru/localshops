package domain

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrCanNotCreateProduct = errors.New("can not create product")
	ErrShopNotFound        = errors.New("shop not found")
)

type ProductServ interface {
	New(
		ctx context.Context,
		id string,
		name string,
		description string,
		shopID string,
		priceAmount int,
		priceCurrency string,
	) (*Product, error)
}

type productServ struct {
	shopServ RemoteShopServ
}

func NewProductServ(shopServ RemoteShopServ) *productServ {
	return &productServ{
		shopServ: shopServ,
	}
}

func (s *productServ) New(
	ctx context.Context,
	id string,
	name string,
	description string,
	shopID string,
	priceAmount int,
	priceCurrency string,
) (*Product, error) {
	exists, err := s.shopServ.ExistsByID(ctx, shopID)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrCanNotCreateProduct, err)
	}

	if !exists {
		return nil, ErrShopNotFound
	}

	return newProduct(
		id,
		name,
		description,
		shopID,
		priceAmount,
		priceCurrency,
	)
}
