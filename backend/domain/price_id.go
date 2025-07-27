package domain

import (
	"errors"

	"github.com/google/uuid"
)

type PriceId struct {
	value string
}

func NewPriceId(value string) (*PriceId, error) {
	if value == "" {
		return nil, errors.New("価格IDは必須です")
	}

	if _, err := uuid.Parse(value); err != nil {
		return nil, errors.New("価格IDの形式が不正です")
	}

	return &PriceId{value: value}, nil
}

func (p *PriceId) Value() string {
	return p.value
}