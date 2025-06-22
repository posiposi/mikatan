package domain

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewUserId(t *testing.T) {
	userID, _ := NewUserID(uuid.NewString())
	if userID.Value() == "" {
		t.Errorf("NewUserId() returned an empty value")
	}
}

func TestUserIdError(t *testing.T) {
	_, err := NewUserID("not uuid value")
	if err == nil {
		t.Errorf("NewUserId() should return an error for invalid UUID")
	}
}

func TestUserIdValue(t *testing.T) {
	value := uuid.NewString()
	userID, _ := NewUserID(value)
	if userID.Value() != value {
		t.Errorf("Value() returned %s, expected %s", userID.Value(), value)
	}
}
