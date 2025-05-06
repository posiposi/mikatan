package domain

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewItem(t *testing.T) {
	userId, _ := NewUserId(uuid.NewString())
	itemName, _ := NewItemName("Test Item")
	stock, _ := NewStock(true)
	description, _ := NewDescription("This is a test item.")

	item, err := NewItem(*userId, *itemName, *stock, *description)

	assert.NoError(t, err, "NewItem() should not return an error")
	assert.NotNil(t, item, "NewItem() should return a valid item")
	assert.Equal(t, *userId, item.userId, "UserId should match")
	assert.Equal(t, *itemName, item.itemName, "ItemName should match")
	assert.Equal(t, *stock, item.stock, "Stock should match")
	assert.Equal(t, *description, item.description, "Description should match")
	assert.WithinDuration(t, time.Now(), item.createdAt, time.Second, "CreatedAt should be close to now")
	assert.WithinDuration(t, time.Now(), item.updatedAt, time.Second, "UpdatedAt should be close to now")
}
