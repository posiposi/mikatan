package domain

import (
	"testing"
)

func TestNewDescription(t *testing.T) {
	value := "test description"
	description, _ := NewDescription(value)
	if description.Value() == "" {
		t.Errorf("NewDescription() returned an empty value")
	}
}

func TestNewDescriptionEmptyError(t *testing.T) {
	_, err := NewDescription("")
	if err == nil {
		t.Errorf("NewDescription() should return an error for empty value")
	}
}

func TestNewDescriptionOverLengthError(t *testing.T) {
	value := string(make([]byte, 192))
	_, err := NewDescription(value)
	if err == nil {
		t.Errorf("NewDescription() should return an error for invalid value")
	}
}

func TestDescriptionValue(t *testing.T) {
	value := "description"
	description, _ := NewDescription(value)
	if description.Value() != value {
		t.Errorf("Value() returned %s, expected %s", description.Value(), value)
	}
}
