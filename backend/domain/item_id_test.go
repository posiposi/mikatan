package domain

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewItemId(t *testing.T) {
	itemID, _ := NewItemID(uuid.NewString())
	if itemID.Value() == "" {
		t.Errorf("NewItemId() returned an empty value")
	}
}

func TestItemIdError(t *testing.T) {
	_, err := NewItemID("not uuid value")
	if err == nil {
		t.Errorf("NewItemId() should return an error for invalid UUID")
	}
}

func TestItemIdValue(t *testing.T) {
	value := uuid.NewString()
	itemID, _ := NewItemID(value)
	if itemID.Value() != value {
		t.Errorf("Value() returned %s, expected %s", itemID.Value(), value)
	}
}
