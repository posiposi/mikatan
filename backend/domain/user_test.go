package domain

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewUser(t *testing.T) {
	userID, _ := NewUserID(uuid.NewString())
	name := "TestUser"
	email, _ := NewEmail("test@example.com")
	password, _ := NewPassword("validPassword123")

	user, err := NewUser(userID, name, email, password)
	if err != nil {
		t.Errorf("NewUser() returned an error: %v", err)
	}

	if user.ID() != userID {
		t.Errorf("ID() returned %v, expected %v", user.ID(), userID)
	}
	if user.Name() != name {
		t.Errorf("Name() returned %s, expected %s", user.Name(), name)
	}
	if !user.Email().Equals(email) {
		t.Errorf("Email() returned %v, expected %v", user.Email(), email)
	}
	if !user.Password().Equals(password) {
		t.Errorf("Password() returned %v, expected %v", user.Password(), password)
	}
}

func TestNewUserWithEmptyNameError(t *testing.T) {
	userID, _ := NewUserID(uuid.NewString())
	name := ""
	email, _ := NewEmail("test@example.com")
	password, _ := NewPassword("validPassword123")

	_, err := NewUser(userID, name, email, password)
	if err == nil {
		t.Errorf("NewUser() should return an error for empty name")
	}
}

func TestNewUserWithNilIDError(t *testing.T) {
	name := "TestUser"
	email, _ := NewEmail("test@example.com")
	password, _ := NewPassword("validPassword123")

	_, err := NewUser(nil, name, email, password)
	if err == nil {
		t.Errorf("NewUser() should return an error for nil ID")
	}
}

func TestNewUserWithNilEmailError(t *testing.T) {
	userID, _ := NewUserID(uuid.NewString())
	name := "TestUser"
	password, _ := NewPassword("validPassword123")

	_, err := NewUser(userID, name, nil, password)
	if err == nil {
		t.Errorf("NewUser() should return an error for nil email")
	}
}

func TestNewUserWithNilPasswordError(t *testing.T) {
	userID, _ := NewUserID(uuid.NewString())
	name := "TestUser"
	email, _ := NewEmail("test@example.com")

	_, err := NewUser(userID, name, email, nil)
	if err == nil {
		t.Errorf("NewUser() should return an error for nil password")
	}
}

func TestUserEquals(t *testing.T) {
	userID1, _ := NewUserID(uuid.NewString())
	userID2, _ := NewUserID(uuid.NewString())
	email1, _ := NewEmail("test1@example.com")
	email2, _ := NewEmail("test2@example.com")
	password, _ := NewPassword("validPassword123")

	user1, _ := NewUser(userID1, "TestUser1", email1, password)
	user2, _ := NewUser(userID1, "TestUser2", email2, password)
	user3, _ := NewUser(userID2, "TestUser3", email1, password)

	if !user1.Equals(user2) {
		t.Errorf("Equals() should return true for users with same ID")
	}

	if user1.Equals(user3) {
		t.Errorf("Equals() should return false for users with different IDs")
	}
}

func TestUserString(t *testing.T) {
	userID, _ := NewUserID(uuid.NewString())
	name := "TestUser"
	email, _ := NewEmail("test@example.com")
	password, _ := NewPassword("validPassword123")

	user, _ := NewUser(userID, name, email, password)
	expected := "User{ID: " + userID.Value() + ", Name: " + name + ", Email: " + email.String() + "}"
	if user.String() != expected {
		t.Errorf("String() returned %s, expected %s", user.String(), expected)
	}
}