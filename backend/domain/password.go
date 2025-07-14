package domain

import (
	"fmt"
	"strings"
)

const (
	MinPasswordLength = 8
	MaxPasswordLength = 128
)

type Password struct {
	value string
}

func NewPassword(value string) (*Password, error) {
	if strings.TrimSpace(value) == "" {
		return nil, fmt.Errorf("password cannot be empty")
	}

	if len(value) < MinPasswordLength {
		return nil, fmt.Errorf("password must be at least %d characters long", MinPasswordLength)
	}

	if len(value) > MaxPasswordLength {
		return nil, fmt.Errorf("password must be less than %d characters long", MaxPasswordLength)
	}

	return &Password{value: value}, nil
}

func (p *Password) Value() string {
	return p.value
}

func (p *Password) String() string {
	return "****"
}

func (p *Password) Equals(other *Password) bool {
	if other == nil {
		return false
	}
	return p.value == other.value
}