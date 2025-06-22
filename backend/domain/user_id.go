package domain

import (
	"fmt"

	"github.com/google/uuid"
)

type UserID struct {
	value string
}

func NewUserID(value string) (*UserID, error) {
	if uuid.Validate(value) != nil {
		return nil, fmt.Errorf("invalid UUID: %s", value)
	}
	userID := new(UserID)
	userID.value = value
	return userID, nil
}

func (userID *UserID) Value() string {
	return userID.value
}
