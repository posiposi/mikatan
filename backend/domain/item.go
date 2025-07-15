package domain

import (
	"time"

	"github.com/google/uuid"
)

type Item struct {
	itemID      ItemID
	userID      UserID
	itemName    ItemName
	stock       Stock
	description Description
	createdAt   time.Time
	updatedAt   time.Time
	deletedAt   time.Time
}

func NewItem(itemID *ItemID, userID UserID, itemName ItemName, stock Stock, description Description) (*Item, error) {
	var id ItemID
	if itemID == nil {
		newID, err := NewItemID(uuid.NewString())
		if err != nil {
			return nil, err
		}
		id = *newID
	} else {
		id = *itemID
	}
	item := &Item{
		itemID:      id,
		userID:      userID,
		itemName:    itemName,
		stock:       stock,
		description: description,
		createdAt:   time.Now(),
		updatedAt:   time.Now(),
	}
	return item, nil
}

func (i *Item) ItemID() string {
	return i.itemID.Value()
}

func (i *Item) UserID() string {
	return i.userID.Value()
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

