package domain

import (
	"fmt"

	"github.com/google/uuid"
)

type ItemID struct {
	value string
}

func NewItemID(value string) (*ItemID, error) {
	if uuid.Validate(value) != nil {
		return nil, fmt.Errorf("invalid UUID: %s", value)
	}
	itemID := new(ItemID)
	itemID.value = value
	return itemID, nil
}

func (itemID *ItemID) Value() string {
	return itemID.value
}
