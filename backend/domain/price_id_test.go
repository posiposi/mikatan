package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPriceId(t *testing.T) {
	t.Run("有効なUUIDの場合、PriceIdが作成される", func(t *testing.T) {
		validUUID := "550e8400-e29b-41d4-a716-446655440000"
		priceId, err := NewPriceId(validUUID)

		assert.NoError(t, err)
		assert.NotNil(t, priceId)
		assert.Equal(t, validUUID, priceId.Value())
	})

	t.Run("無効なUUIDの場合、エラーが返される", func(t *testing.T) {
		invalidUUID := "invalid-uuid"
		priceId, err := NewPriceId(invalidUUID)

		assert.Error(t, err)
		assert.Nil(t, priceId)
	})

	t.Run("空文字列の場合、エラーが返される", func(t *testing.T) {
		emptyString := ""
		priceId, err := NewPriceId(emptyString)

		assert.Error(t, err)
		assert.Nil(t, priceId)
	})
}