package domain

import "errors"

type PriceWithTax struct {
	value int
}

const (
	minPrice = 0
	maxPrice = 100000000
)

func NewPriceWithTax(value int) (*PriceWithTax, error) {
	if value < minPrice {
		return nil, errors.New("税込価格は0以上である必要があります")
	}

	if value >= maxPrice {
		return nil, errors.New("税込価格は1億円未満である必要があります")
	}

	return &PriceWithTax{value: value}, nil
}

func (p *PriceWithTax) Value() int {
	return p.value
}