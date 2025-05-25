package domain

import (
	"fmt"

	"github.com/google/uuid"
)

type UserId struct {
	value string
}

func NewUserId(value string) (*UserId, error) {
	if uuid.Validate(value) != nil {
		return nil, fmt.Errorf("invalid UUID: %s", value)
	}
	userId := new(UserId)
	userId.value = value
	return userId, nil
}

func (userId *UserId) Value() string {
	return userId.value
}
