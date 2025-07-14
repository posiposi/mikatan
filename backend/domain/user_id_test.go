package domain

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewUserId(t *testing.T) {
	userId, _ := NewUserId(uuid.NewString())
	if userId.Value() == "" {
		t.Errorf("NewUserId() returned an empty value")
	}
}

func TestUserIdError(t *testing.T) {
	_, err := NewUserId("not uuid value")
	if err == nil {
		t.Errorf("NewUserId() should return an error for invalid UUID")
	}
}

func TestUserIdValue(t *testing.T) {
	value := uuid.NewString()
	userId, _ := NewUserId(value)
	if userId.Value() != value {
		t.Errorf("Value() returned %s, expected %s", userId.Value(), value)
	}
}
