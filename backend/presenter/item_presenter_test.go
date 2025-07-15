package presenter

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/posiposi/project/backend/domain"
	"github.com/stretchr/testify/assert"
)

func createTestDomainItem() *domain.Item {
	userID, _ := domain.NewUserID(uuid.NewString())
	itemName, _ := domain.NewItemName("Test Item")
	stock, _ := domain.NewStock(true)
	description, _ := domain.NewDescription("Test Description")
	
	item, _ := domain.NewItem(nil, *userID, *itemName, *stock, *description)
	return item
}

func TestItemPresenter_ToJSON(t *testing.T) {
	presenter := NewItemPresenter()
	domainItem := createTestDomainItem()
	
	result := presenter.ToJSON(domainItem)
	
	assert.Equal(t, domainItem.ItemID(), result.ItemId)
	assert.Equal(t, domainItem.UserID(), result.UserId)
	assert.Equal(t, domainItem.ItemName(), result.ItemName)
	assert.Equal(t, domainItem.Stock(), result.Stock)
	assert.Equal(t, domainItem.Description(), result.Description)
	assert.IsType(t, time.Time{}, result.CreatedAt)
	assert.IsType(t, time.Time{}, result.UpdatedAt)
}

func TestItemPresenter_ToJSONList(t *testing.T) {
	presenter := NewItemPresenter()
	domainItems := []*domain.Item{
		createTestDomainItem(),
		createTestDomainItem(),
	}
	
	result := presenter.ToJSONList(domainItems)
	
	assert.Len(t, result, 2)
	
	for i, jsonItem := range result {
		assert.Equal(t, domainItems[i].ItemID(), jsonItem.ItemId)
		assert.Equal(t, domainItems[i].UserID(), jsonItem.UserId)
		assert.Equal(t, domainItems[i].ItemName(), jsonItem.ItemName)
		assert.Equal(t, domainItems[i].Stock(), jsonItem.Stock)
		assert.Equal(t, domainItems[i].Description(), jsonItem.Description)
	}
}

func TestItemPresenter_ToJSONList_EmptySlice(t *testing.T) {
	presenter := NewItemPresenter()
	emptyItems := []*domain.Item{}
	
	result := presenter.ToJSONList(emptyItems)
	
	assert.Empty(t, result)
	assert.NotNil(t, result)
}