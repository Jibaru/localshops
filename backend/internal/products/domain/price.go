package domain

import "errors"

var (
	ErrPriceAmountCanNotBeNegative = errors.New("price amount can not be negative")
)

type Price struct {
	currency *Currency
	amount   int
}

func NewPrice(currency string, amount int) (*Price, error) {
	if amount < 0 {
		return nil, ErrPriceAmountCanNotBeNegative
	}

	currencyVal, err := NewCurrency(currency)
	if err != nil {
		return nil, err
	}

	return HydratePrice(currencyVal, amount), nil
}

func HydratePrice(currency *Currency, amount int) *Price {
	return &Price{
		currency: currency,
		amount:   amount,
	}
}

func (p *Price) Currency() string {
	return p.currency.Value()
}

func (p *Price) Amount() int {
	return p.amount
}
