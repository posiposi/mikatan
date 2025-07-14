package domain

import (
	"testing"
)

func TestNewPassword(t *testing.T) {
	value := "validPassword123"
	password, err := NewPassword(value)
	if err != nil {
		t.Errorf("NewPassword() returned an error: %v", err)
	}
	if password.Value() != value {
		t.Errorf("Value() returned %s, expected %s", password.Value(), value)
	}
}

func TestNewPasswordEmptyError(t *testing.T) {
	_, err := NewPassword("")
	if err == nil {
		t.Errorf("NewPassword() should return an error for empty value")
	}
}

func TestNewPasswordTooShortError(t *testing.T) {
	shortPassword := "short"
	_, err := NewPassword(shortPassword)
	if err == nil {
		t.Errorf("NewPassword() should return an error for password shorter than %d characters", MinPasswordLength)
	}
}

func TestNewPasswordTooLongError(t *testing.T) {
	longPassword := string(make([]byte, MaxPasswordLength+1))
	_, err := NewPassword(longPassword)
	if err == nil {
		t.Errorf("NewPassword() should return an error for password longer than %d characters", MaxPasswordLength)
	}
}

func TestPasswordValue(t *testing.T) {
	value := "validPassword123"
	password, _ := NewPassword(value)
	if password.Value() != value {
		t.Errorf("Value() returned %s, expected %s", password.Value(), value)
	}
}

func TestPasswordString(t *testing.T) {
	value := "validPassword123"
	password, _ := NewPassword(value)
	expected := "****"
	if password.String() != expected {
		t.Errorf("String() returned %s, expected %s", password.String(), expected)
	}
}

func TestPasswordEquals(t *testing.T) {
	value := "validPassword123"
	password1, _ := NewPassword(value)
	password2, _ := NewPassword(value)
	differentPassword, _ := NewPassword("differentPassword123")

	if !password1.Equals(password2) {
		t.Errorf("Equals() should return true for same password values")
	}

	if password1.Equals(differentPassword) {
		t.Errorf("Equals() should return false for different password values")
	}
}