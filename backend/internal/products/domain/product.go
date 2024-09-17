package domain

import (
	"errors"
	"strings"
)

var (
	ErrEmptyID          = errors.New("empty id")
	ErrEmptyName        = errors.New("empty name")
	ErrEmptyDescription = errors.New("empty description")
	ErrEmptyShopID      = errors.New("empty shop id")
)

type Product struct {
	id          string
	name        string
	description string
	price       *Price
	shopID      string
}

func newProduct(
	id string,
	name string,
	description string,
	shopID string,
	priceAmount int,
	priceCurrency string,
) (*Product, error) {
	if strings.TrimSpace(id) == "" {
		return nil, ErrEmptyID
	}

	if strings.TrimSpace(name) == "" {
		return nil, ErrEmptyName
	}

	if strings.TrimSpace(description) == "" {
		return nil, ErrEmptyDescription
	}

	if strings.TrimSpace(shopID) == "" {
		return nil, ErrEmptyShopID
	}

	price, err := NewPrice(priceCurrency, priceAmount)
	if err != nil {
		return nil, err
	}

	return HydrateProduct(id, name, description, shopID, price), nil
}

func HydrateProduct(
	id string,
	name string,
	description string,
	shopID string,
	price *Price,
) *Product {
	return &Product{
		id:          id,
		name:        name,
		description: description,
		shopID:      shopID,
		price:       price,
	}
}

func (p *Product) ID() string {
	return p.id
}

func (p *Product) Name() string {
	return p.name
}

func (p *Product) Description() string {
	return p.description
}

func (p *Product) ShopID() string {
	return p.shopID
}

func (p *Product) PriceAmount() int {
	return p.price.Amount()
}

func (p *Product) PriceCurrency() string {
	return p.price.Currency()
}
