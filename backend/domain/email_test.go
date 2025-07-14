package domain

import (
	"testing"
)

func TestNewEmail(t *testing.T) {
	value := "test@example.com"
	email, err := NewEmail(value)
	if err != nil {
		t.Errorf("NewEmail() returned an error: %v", err)
	}
	if email.Value() != value {
		t.Errorf("Value() returned %s, expected %s", email.Value(), value)
	}
}

func TestNewEmailEmptyError(t *testing.T) {
	_, err := NewEmail("")
	if err == nil {
		t.Errorf("NewEmail() should return an error for empty value")
	}
}

func TestNewEmailInvalidFormatError(t *testing.T) {
	testCases := []string{
		"invalid-email",
		"@example.com",
		"test@",
		"test@.com",
		"test..test@example.com",
		"test@example..com",
	}

	for _, testCase := range testCases {
		t.Run(testCase, func(t *testing.T) {
			_, err := NewEmail(testCase)
			if err == nil {
				t.Errorf("NewEmail() should return an error for invalid email format: %s", testCase)
			}
		})
	}
}

func TestEmailValue(t *testing.T) {
	value := "valid@example.com"
	email, _ := NewEmail(value)
	if email.Value() != value {
		t.Errorf("Value() returned %s, expected %s", email.Value(), value)
	}
}

func TestEmailString(t *testing.T) {
	value := "valid@example.com"
	email, _ := NewEmail(value)
	if email.String() != value {
		t.Errorf("String() returned %s, expected %s", email.String(), value)
	}
}

func TestEmailEquals(t *testing.T) {
	value := "test@example.com"
	email1, _ := NewEmail(value)
	email2, _ := NewEmail(value)
	differentEmail, _ := NewEmail("different@example.com")

	if !email1.Equals(email2) {
		t.Errorf("Equals() should return true for same email values")
	}

	if email1.Equals(differentEmail) {
		t.Errorf("Equals() should return false for different email values")
	}
}