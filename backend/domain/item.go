package domain

import (
	"time"

	"github.com/google/uuid"
)

type Item struct {
	itemId      ItemId
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
		itemId:      id,
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
	return i.itemId.Value()
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

func (i *Item) CreatedAt() time.Time {
	return i.createdAt
}

func (i *Item) UpdatedAt() time.Time {
	return i.updatedAt
}

