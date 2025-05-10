package domain

import (
	"fmt"
)

type ItemName struct {
	value string
}

func NewItemName(value string) (*ItemName, error) {
	if len(value) == 0 {
		return nil, fmt.Errorf("value count must be greater than 0")
	}

	if len(value) > 191 {
		return nil, fmt.Errorf("value count must be less than 191")
	}

	itemName := new(ItemName)
	itemName.value = value
	return itemName, nil
}

func (itemName *ItemName) Value() string {
	return itemName.value
}
