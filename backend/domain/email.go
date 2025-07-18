package domain

import (
	"fmt"
	"regexp"
	"strings"
)

type Email struct {
	value string
}

func NewEmail(value string) (*Email, error) {
	if strings.TrimSpace(value) == "" {
		return nil, fmt.Errorf("email cannot be empty")
	}

	if !isValidEmail(value) {
		return nil, fmt.Errorf("invalid email format: %s", value)
	}

	return &Email{value: value}, nil
}

func (e *Email) Value() string {
	return e.value
}

func (e *Email) String() string {
	return e.value
}

func (e *Email) Equals(other *Email) bool {
	if other == nil {
		return false
	}
	return e.value == other.value
}

func isValidEmail(email string) bool {
	if strings.Contains(email, "..") {
		return false
	}
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9]([a-zA-Z0-9._-]*[a-zA-Z0-9])?@[a-zA-Z0-9]([a-zA-Z0-9.-]*[a-zA-Z0-9])?\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}