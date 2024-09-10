package domain

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidCurrency = errors.New("invalid currency")
)

type Currency struct {
	value string
}

func NewCurrency(value string) (*Currency, error) {
	if value != "PEN" {
		return nil, fmt.Errorf("%w: %v", ErrInvalidCurrency, value)
	}

	return HydrateCurrency(value), nil
}

func HydrateCurrency(value string) *Currency {
	return &Currency{
		value: value,
	}
}

func (c *Currency) Value() string {
	return c.value
}
