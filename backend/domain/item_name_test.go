package domain

import (
	"testing"
)

func TestNewItemName(t *testing.T) {
	value := "test item name"
	itemName, _ := NewItemName(value)
	if itemName.Value() == "" {
		t.Errorf("NewItemNameId() returned an empty value")
	}
}

func TestNewItemNameEmptyError(t *testing.T) {
	_, err := NewItemName("")
	if err == nil {
		t.Errorf("NewItemName() should return an error for empty value")
	}
}

func TestNewItemNameOverLengthError(t *testing.T) {
	value := string(make([]byte, 192))
	_, err := NewItemName(value)
	if err == nil {
		t.Errorf("NewItemName() should return an error for invalid value")
	}
}

func TestItemNameValue(t *testing.T) {
	value := "test item name"
	itemName, _ := NewItemName(value)
	if itemName.Value() != value {
		t.Errorf("Value() returned %s, expected %s", itemName.Value(), value)
	}
}
