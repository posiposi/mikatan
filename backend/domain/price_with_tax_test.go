package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPriceWithTax(t *testing.T) {
	t.Run("有効な金額の場合、PriceWithTaxが作成される", func(t *testing.T) {
		validAmounts := []int{0, 100, 1000, 99999999}
		for _, amount := range validAmounts {
			priceWithTax, err := NewPriceWithTax(amount)

			assert.NoError(t, err)
			assert.NotNil(t, priceWithTax)
			assert.Equal(t, amount, priceWithTax.Value())
		}
	})

	t.Run("負の金額の場合、エラーが返される", func(t *testing.T) {
		negativeAmount := -100
		priceWithTax, err := NewPriceWithTax(negativeAmount)

		assert.Error(t, err)
		assert.Nil(t, priceWithTax)
	})

	t.Run("最大値を超える金額の場合、エラーが返される", func(t *testing.T) {
		overAmount := 100000000
		priceWithTax, err := NewPriceWithTax(overAmount)

		assert.Error(t, err)
		assert.Nil(t, priceWithTax)
	})
}