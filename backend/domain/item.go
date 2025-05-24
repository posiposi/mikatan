package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/posiposi/project/backend/dto"
)

type Item struct {
	itemID      ItemId
	userId      UserId
	itemName    ItemName
	stock       Stock
	description Description
	createdAt   time.Time
	updatedAt   time.Time
	deletedAt   time.Time
}

func NewItem(itemId *ItemId, userId UserId, itemName ItemName, stock Stock, description Description) (*Item, error) {
	var id ItemId
	if itemId == nil {
		newId, err := NewItemId(uuid.NewString())
		if err != nil {
			return nil, err
		}
		id = *newId
	} else {
		id = *itemId
	}
	item := &Item{
		itemID:      id,
		userId:      userId,
		itemName:    itemName,
		stock:       stock,
		description: description,
		createdAt:   time.Now(),
		updatedAt:   time.Now(),
	}
	return item, nil
}

func (i *Item) ItemId() string {
	return i.itemID.Value()
}

func (i *Item) UserId() string {
	return i.userId.Value()
}

func (i *Item) ItemName() string {
	return i.itemName.Value()
}

func (i *Item) Stock() bool {
	return i.stock.Value()
}

func (i *Item) Description() string {
	return i.description.Value()
}

func (i *Item) ToDto() dto.ItemResponse {
	return dto.ItemResponse{
		ItemId:      i.itemID.Value(),
		UserId:      i.userId.Value(),
		ItemName:    i.itemName.Value(),
		Stock:       i.stock.Value(),
		Description: i.description.Value(),
		CreatedAt:   i.createdAt,
		UpdatedAt:   i.updatedAt,
	}
}
