package domain

import "errors"

type TaxRate struct {
	value float64
}

const allowedTaxRate = 10.0

func NewTaxRate(value float64) (*TaxRate, error) {
	if value != allowedTaxRate {
		return nil, errors.New("税率は10%のみ使用可能です")
	}

	return &TaxRate{value: value}, nil
}

func (t *TaxRate) Value() float64 {
	return t.value
}

func (t *TaxRate) Multiplier() float64 {
	const percentToDecimal = 100
	return t.value / percentToDecimal
}