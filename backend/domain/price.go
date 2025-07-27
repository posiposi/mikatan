package domain

import (
	"time"

	"github.com/google/uuid"
)

type Price struct {
	priceId         PriceId
	itemId          ItemId
	priceWithTax    PriceWithTax
	priceWithoutTax PriceWithoutTax
	taxRate         TaxRate
	currency        Currency
	startDate       time.Time
	endDate         *time.Time
	createdAt       time.Time
	updatedAt       time.Time
}

func NewPrice(
	priceId *PriceId,
	itemId ItemId,
	priceWithTax PriceWithTax,
	priceWithoutTax PriceWithoutTax,
	taxRate TaxRate,
	currency Currency,
	startDate time.Time,
	endDate *time.Time,
) (*Price, error) {
	var id PriceId
	if priceId == nil {
		newId, err := NewPriceId(uuid.NewString())
		if err != nil {
			return nil, err
		}
		id = *newId
	} else {
		id = *priceId
	}

	price := &Price{
		priceId:         id,
		itemId:          itemId,
		priceWithTax:    priceWithTax,
		priceWithoutTax: priceWithoutTax,
		taxRate:         taxRate,
		currency:        currency,
		startDate:       startDate,
		endDate:         endDate,
		createdAt:       time.Now(),
		updatedAt:       time.Now(),
	}

	return price, nil
}

func (p *Price) PriceId() string {
	return p.priceId.Value()
}

func (p *Price) ItemId() string {
	return p.itemId.Value()
}

func (p *Price) PriceWithTax() int {
	return p.priceWithTax.Value()
}

func (p *Price) PriceWithoutTax() int {
	return p.priceWithoutTax.Value()
}

func (p *Price) TaxRate() float64 {
	return p.taxRate.Value()
}

func (p *Price) Currency() string {
	return p.currency.Value()
}

func (p *Price) StartDate() time.Time {
	return p.startDate
}

func (p *Price) EndDate() *time.Time {
	return p.endDate
}

func (p *Price) CreatedAt() time.Time {
	return p.createdAt
}

func (p *Price) UpdatedAt() time.Time {
	return p.updatedAt
}