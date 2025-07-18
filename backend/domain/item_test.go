package domain

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func createTestItem() (*Item, ItemId, error) {
	userId, _ := NewUserId(uuid.NewString())
	itemName, _ := NewItemName("Test Item")
	stock, _ := NewStock(true)
	description, _ := NewDescription("This is a test item.")

	item, err := NewItem(nil, *userId, *itemName, *stock, *description)
	return item, item.itemId, err
}

func TestNewItem(t *testing.T) {
	item, _, err := createTestItem()
	if err != nil {
		t.Fatalf("Failed to create test item: %v", err)
	}
	assert.NotNil(t, item, "NewItem() should return a non-nil Item")
	assert.NotNil(t, item.itemId, "ItemId should not be nil")
	assert.NotNil(t, item.userId, "UserId should not be nil")
	assert.NotNil(t, item.itemName, "ItemName should not be nil")
	assert.NotNil(t, item.stock, "Stock should not be nil")
	assert.NotNil(t, item.description, "Description should not be nil")
}

func TestItemId(t *testing.T) {
	item, itemId, _ := createTestItem()
	assert.Equal(t, itemId.Value(), item.ItemId())
}

func TestUserId(t *testing.T) {
	item, _, _ := createTestItem()
	assert.Equal(t, item.userId.Value(), item.UserId())
}

func TestItemName(t *testing.T) {
	item, _, _ := createTestItem()
	assert.Equal(t, item.itemName.Value(), item.ItemName())
}

func TestStock(t *testing.T) {
	item, _, _ := createTestItem()
	assert.Equal(t, item.stock.Value(), item.Stock())
}

func TestDescription(t *testing.T) {
	item, _, _ := createTestItem()
	assert.Equal(t, item.description.Value(), item.Description())
}
