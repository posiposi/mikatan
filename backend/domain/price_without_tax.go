package domain

import "errors"

type PriceWithoutTax struct {
	value int
}

const (
	minPriceWithoutTax = 0
	maxPriceWithoutTax = 100000000
)

func NewPriceWithoutTax(value int) (*PriceWithoutTax, error) {
	if value < minPriceWithoutTax {
		return nil, errors.New("税抜価格は0以上である必要があります")
	}

	if value >= maxPriceWithoutTax {
		return nil, errors.New("税抜価格は1億円未満である必要があります")
	}

	return &PriceWithoutTax{value: value}, nil
}

func (p *PriceWithoutTax) Value() int {
	return p.value
}