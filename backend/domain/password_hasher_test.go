package domain

import (
	"testing"
)

func TestNewPasswordHasher(t *testing.T) {
	hasher := NewPasswordHasher()
	if hasher == nil {
		t.Errorf("NewPasswordHasher() should not return nil")
	}
}

func TestPasswordHasherHash(t *testing.T) {
	hasher := NewPasswordHasher()
	password, _ := NewPassword("validPassword123")

	hashedPassword, err := hasher.Hash(password)
	if err != nil {
		t.Errorf("Hash() returned an error: %v", err)
	}

	if hashedPassword == nil {
		t.Errorf("Hash() should not return nil")
	}

	if hashedPassword.Value() == password.Value() {
		t.Errorf("Hash() should return a different value than the original password")
	}

	if len(hashedPassword.Value()) == 0 {
		t.Errorf("Hash() should return a non-empty hashed password")
	}
}

func TestPasswordHasherHashNilPassword(t *testing.T) {
	hasher := NewPasswordHasher()

	_, err := hasher.Hash(nil)
	if err == nil {
		t.Errorf("Hash() should return an error for nil password")
	}
}

func TestPasswordHasherVerify(t *testing.T) {
	hasher := NewPasswordHasher()
	password, _ := NewPassword("validPassword123")

	hashedPassword, _ := hasher.Hash(password)

	isValid := hasher.Verify(password, hashedPassword)
	if !isValid {
		t.Errorf("Verify() should return true for correct password")
	}

	wrongPassword, _ := NewPassword("wrongPassword123")
	isValid = hasher.Verify(wrongPassword, hashedPassword)
	if isValid {
		t.Errorf("Verify() should return false for incorrect password")
	}
}

func TestPasswordHasherVerifyNilPassword(t *testing.T) {
	hasher := NewPasswordHasher()
	password, _ := NewPassword("validPassword123")
	hashedPassword, _ := hasher.Hash(password)

	isValid := hasher.Verify(nil, hashedPassword)
	if isValid {
		t.Errorf("Verify() should return false for nil password")
	}
}

func TestPasswordHasherVerifyNilHashedPassword(t *testing.T) {
	hasher := NewPasswordHasher()
	password, _ := NewPassword("validPassword123")

	isValid := hasher.Verify(password, nil)
	if isValid {
		t.Errorf("Verify() should return false for nil hashed password")
	}
}

func TestPasswordHasherConsistency(t *testing.T) {
	hasher := NewPasswordHasher()
	password, _ := NewPassword("validPassword123")

	hashedPassword1, _ := hasher.Hash(password)
	hashedPassword2, _ := hasher.Hash(password)

	if hashedPassword1.Value() == hashedPassword2.Value() {
		t.Errorf("Hash() should return different values for the same password (salt should be different)")
	}

	if !hasher.Verify(password, hashedPassword1) {
		t.Errorf("Verify() should return true for hashedPassword1")
	}

	if !hasher.Verify(password, hashedPassword2) {
		t.Errorf("Verify() should return true for hashedPassword2")
	}
}