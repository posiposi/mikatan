package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewPrice(t *testing.T) {
	t.Run("有効な値でPriceが作成される", func(t *testing.T) {
		priceId, _ := NewPriceId("550e8400-e29b-41d4-a716-446655440000")
		itemId, _ := NewItemId("550e8400-e29b-41d4-a716-446655440001")
		priceWithTax, _ := NewPriceWithTax(1100)
		priceWithoutTax, _ := NewPriceWithoutTax(1000)
		currency, _ := NewCurrency("JPY")
		taxRate, _ := NewTaxRate(10.0)
		startDate := time.Now()
		endDate := time.Now().Add(24 * time.Hour)

		price, err := NewPrice(
			priceId,
			*itemId,
			*priceWithTax,
			*priceWithoutTax,
			*taxRate,
			*currency,
			startDate,
			&endDate,
		)

		assert.NoError(t, err)
		assert.NotNil(t, price)
		assert.Equal(t, "550e8400-e29b-41d4-a716-446655440000", price.PriceId())
		assert.Equal(t, "550e8400-e29b-41d4-a716-446655440001", price.ItemId())
		assert.Equal(t, 1100, price.PriceWithTax())
		assert.Equal(t, 1000, price.PriceWithoutTax())
		assert.Equal(t, 10.0, price.TaxRate())
		assert.Equal(t, "JPY", price.Currency())
	})

	t.Run("nilのpriceIdでもPriceが作成される", func(t *testing.T) {
		itemId, _ := NewItemId("550e8400-e29b-41d4-a716-446655440001")
		priceWithTax, _ := NewPriceWithTax(1100)
		priceWithoutTax, _ := NewPriceWithoutTax(1000)
		currency, _ := NewCurrency("JPY")
		taxRate, _ := NewTaxRate(10.0)
		startDate := time.Now()

		price, err := NewPrice(
			nil,
			*itemId,
			*priceWithTax,
			*priceWithoutTax,
			*taxRate,
			*currency,
			startDate,
			nil,
		)

		assert.NoError(t, err)
		assert.NotNil(t, price)
		assert.NotEmpty(t, price.PriceId())
	})
}