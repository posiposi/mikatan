package domain

import (
	"errors"
	"strings"
)

type Currency struct {
	value string
}

var validCurrencies = map[string]bool{
	"JPY": true,
	"USD": true,
	"EUR": true,
	"GBP": true,
	"CNY": true,
	"KRW": true,
}

func NewCurrency(value string) (*Currency, error) {
	if value == "" {
		return nil, errors.New("通貨コードは必須です")
	}

	upperValue := strings.ToUpper(value)
	const currencyCodeLength = 3
	if len(upperValue) != currencyCodeLength {
		return nil, errors.New("通貨コードは3文字である必要があります")
	}

	if !validCurrencies[upperValue] {
		return nil, errors.New("サポートされていない通貨コードです")
	}

	return &Currency{value: upperValue}, nil
}

func (c *Currency) Value() string {
	return c.value
}