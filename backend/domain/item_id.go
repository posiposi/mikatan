package domain

import (
	"fmt"

	"github.com/google/uuid"
)

type ItemId struct {
	value string
}

func NewItemId(value string) (*ItemId, error) {
	if uuid.Validate(value) != nil {
		return nil, fmt.Errorf("invalid UUID: %s", value)
	}
	itemId := new(ItemId)
	itemId.value = value
	return itemId, nil
}

func (itemId *ItemId) Value() string {
	return itemId.value
}
