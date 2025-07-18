package domain

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type PasswordHasher struct {
	cost int
}

func NewPasswordHasher() *PasswordHasher {
	return &PasswordHasher{
		cost: bcrypt.DefaultCost,
	}
}

func (ph *PasswordHasher) Hash(password *Password) (*Password, error) {
	if password == nil {
		return nil, fmt.Errorf("password cannot be nil")
	}

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password.Value()), ph.cost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	hashedPassword := &Password{value: string(hashedBytes)}
	return hashedPassword, nil
}

func (ph *PasswordHasher) Verify(password *Password, hashedPassword *Password) bool {
	if password == nil || hashedPassword == nil {
		return false
	}

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword.Value()), []byte(password.Value()))
	return err == nil
}