package domain

import (
	"errors"
	"strings"
)

var (
	ErrEmptyID          = errors.New("empty id")
	ErrEmptyName        = errors.New("empty name")
	ErrEmptyDescription = errors.New("empty description")
)

type Product struct {
	id          string
	name        string
	description string
	price       *Price
}

func NewProduct(
	id string,
	name string,
	description string,
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

	price, err := NewPrice(priceCurrency, priceAmount)
	if err != nil {
		return nil, err
	}

	return HydrateProduct(id, name, description, price), nil
}

func HydrateProduct(
	id string,
	name string,
	description string,
	price *Price,
) *Product {
	return &Product{
		id:          id,
		name:        name,
		description: description,
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

func (p *Product) PriceAmount() int {
	return p.price.Amount()
}

func (p *Product) PriceCurrency() string {
	return p.price.Currency()
}
