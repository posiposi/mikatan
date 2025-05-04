package domain

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewItemId(t *testing.T) {
	itemId, _ := NewItemId(uuid.NewString())
	if itemId.Value() == "" {
		t.Errorf("NewItemId() returned an empty value")
	}
}

func TestItemIdError(t *testing.T) {
	_, err := NewItemId("not uuid value")
	if err == nil {
		t.Errorf("NewItemId() should return an error for invalid UUID")
	}
}

func TestValue(t *testing.T) {
	value := uuid.NewString()
	itemId, _ := NewItemId(value)
	if itemId.Value() != value {
		t.Errorf("Value() returned %s, expected %s", itemId.Value(), value)
	}
}
