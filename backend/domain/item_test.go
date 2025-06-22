package domain

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func createTestItem() (*Item, ItemID, error) {
	userID, _ := NewUserID(uuid.NewString())
	itemName, _ := NewItemName("Test Item")
	stock, _ := NewStock(true)
	description, _ := NewDescription("This is a test item.")

	item, err := NewItem(nil, *userID, *itemName, *stock, *description)
	return item, item.itemID, err
}

func TestNewItem(t *testing.T) {
	item, _, err := createTestItem()
	if err != nil {
		t.Fatalf("Failed to create test item: %v", err)
	}
	assert.NotNil(t, item, "NewItem() should return a non-nil Item")
	assert.NotNil(t, item.itemID, "ItemId should not be nil")
	assert.NotNil(t, item.userID, "UserId should not be nil")
	assert.NotNil(t, item.itemName, "ItemName should not be nil")
	assert.NotNil(t, item.stock, "Stock should not be nil")
	assert.NotNil(t, item.description, "Description should not be nil")
}

func TestItemId(t *testing.T) {
	item, itemID, _ := createTestItem()
	assert.Equal(t, itemID.Value(), item.ItemID())
}

func TestUserId(t *testing.T) {
	item, _, _ := createTestItem()
	assert.Equal(t, item.userID.Value(), item.UserID())
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
