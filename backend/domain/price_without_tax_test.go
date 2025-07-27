package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPriceWithoutTax(t *testing.T) {
	t.Run("有効な金額の場合、PriceWithoutTaxが作成される", func(t *testing.T) {
		validAmounts := []int{0, 100, 1000, 99999999}
		for _, amount := range validAmounts {
			priceWithoutTax, err := NewPriceWithoutTax(amount)

			assert.NoError(t, err)
			assert.NotNil(t, priceWithoutTax)
			assert.Equal(t, amount, priceWithoutTax.Value())
		}
	})

	t.Run("負の金額の場合、エラーが返される", func(t *testing.T) {
		negativeAmount := -100
		priceWithoutTax, err := NewPriceWithoutTax(negativeAmount)

		assert.Error(t, err)
		assert.Nil(t, priceWithoutTax)
	})

	t.Run("最大値を超える金額の場合、エラーが返される", func(t *testing.T) {
		overAmount := 100000000
		priceWithoutTax, err := NewPriceWithoutTax(overAmount)

		assert.Error(t, err)
		assert.Nil(t, priceWithoutTax)
	})
}