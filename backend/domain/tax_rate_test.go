package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTaxRate(t *testing.T) {
	t.Run("有効な税率の場合、TaxRateが作成される", func(t *testing.T) {
		validRate := 10.0
		taxRate, err := NewTaxRate(validRate)

		assert.NoError(t, err)
		assert.NotNil(t, taxRate)
		assert.Equal(t, validRate, taxRate.Value())
	})

	t.Run("負の税率の場合、エラーが返される", func(t *testing.T) {
		negativeRate := -5.0
		taxRate, err := NewTaxRate(negativeRate)

		assert.Error(t, err)
		assert.Nil(t, taxRate)
	})

	t.Run("10.0以外の税率の場合、エラーが返される", func(t *testing.T) {
		invalidRates := []float64{0, 5, 8, 10.5, 101.0}
		for _, rate := range invalidRates {
			taxRate, err := NewTaxRate(rate)

			assert.Error(t, err)
			assert.Nil(t, taxRate)
		}
	})

	t.Run("計算用の乗数を取得できる", func(t *testing.T) {
		taxRate, _ := NewTaxRate(10.0)

		assert.Equal(t, 0.1, taxRate.Multiplier())
	})
}
