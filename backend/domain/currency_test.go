package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCurrency(t *testing.T) {
	t.Run("有効な通貨コードの場合、Currencyが作成される", func(t *testing.T) {
		validCodes := []string{"JPY", "USD", "EUR"}
		for _, code := range validCodes {
			currency, err := NewCurrency(code)

			assert.NoError(t, err)
			assert.NotNil(t, currency)
			assert.Equal(t, code, currency.Value())
		}
	})

	t.Run("無効な通貨コードの場合、エラーが返される", func(t *testing.T) {
		invalidCodes := []string{"JP", "JPYY", "123", ""}
		for _, code := range invalidCodes {
			currency, err := NewCurrency(code)

			assert.Error(t, err)
			assert.Nil(t, currency)
		}
	})

	t.Run("小文字の通貨コードの場合、大文字に変換される", func(t *testing.T) {
		currency, err := NewCurrency("jpy")

		assert.NoError(t, err)
		assert.NotNil(t, currency)
		assert.Equal(t, "JPY", currency.Value())
	})
}