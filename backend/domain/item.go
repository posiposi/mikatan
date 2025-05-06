package domain

import (
	"time"

	"github.com/google/uuid"
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

func NewItem(userId UserId, itemName ItemName, stock Stock, description Description) (*Item, error) {
	itemId, err := NewItemId(uuid.NewString())
	if err != nil {
		return nil, err
	}
	item := &Item{
		itemID:      *itemId,
		userId:      userId,
		itemName:    itemName,
		stock:       stock,
		description: description,
		createdAt:   time.Now(),
		updatedAt:   time.Now(),
	}
	return item, nil
}
