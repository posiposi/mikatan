package domain

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewUser(t *testing.T) {
	userId, _ := NewUserId(uuid.NewString())
	name := "TestUser"
	email, _ := NewEmail("test@example.com")
	password, _ := NewPassword("validPassword123")

	user, err := NewUser(userId, name, email, password)
	if err != nil {
		t.Errorf("NewUser() returned an error: %v", err)
	}

	if user.Id() != userId {
		t.Errorf("Id() returned %v, expected %v", user.Id(), userId)
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
	if user.Role() == nil {
		t.Errorf("Role() returned nil, expected default USER role")
	}
	if user.Role().Value() != "USER" {
		t.Errorf("Role() returned %s, expected USER", user.Role().Value())
	}
}

func TestNewUserWithRole(t *testing.T) {
	userId, _ := NewUserId(uuid.NewString())
	name := "Test User"
	email, _ := NewEmail("test@example.com")
	password, _ := NewPassword("password123")
	role, _ := NewRole("ADMINISTRATOR")

	user, err := NewUserWithRole(userId, name, email, password, role)
	if err != nil {
		t.Errorf("NewUserWithRole() returned an error: %v", err)
	}

	if user.Role() == nil {
		t.Errorf("Role() returned nil")
	}
	if user.Role().Value() != "ADMINISTRATOR" {
		t.Errorf("Role() returned %s, expected ADMINISTRATOR", user.Role().Value())
	}
}

func TestNewUserWithEmptyNameError(t *testing.T) {
	userId, _ := NewUserId(uuid.NewString())
	name := ""
	email, _ := NewEmail("test@example.com")
	password, _ := NewPassword("validPassword123")

	_, err := NewUser(userId, name, email, password)
	if err == nil {
		t.Errorf("NewUser() should return an error for empty name")
	}
}

func TestNewUserWithNilIdError(t *testing.T) {
	name := "TestUser"
	email, _ := NewEmail("test@example.com")
	password, _ := NewPassword("validPassword123")

	_, err := NewUser(nil, name, email, password)
	if err == nil {
		t.Errorf("NewUser() should return an error for nil Id")
	}
}

func TestNewUserWithNilEmailError(t *testing.T) {
	userId, _ := NewUserId(uuid.NewString())
	name := "TestUser"
	password, _ := NewPassword("validPassword123")

	_, err := NewUser(userId, name, nil, password)
	if err == nil {
		t.Errorf("NewUser() should return an error for nil email")
	}
}

func TestNewUserWithNilPasswordError(t *testing.T) {
	userId, _ := NewUserId(uuid.NewString())
	name := "TestUser"
	email, _ := NewEmail("test@example.com")

	_, err := NewUser(userId, name, email, nil)
	if err == nil {
		t.Errorf("NewUser() should return an error for nil password")
	}
}

func TestUserEquals(t *testing.T) {
	userId1, _ := NewUserId(uuid.NewString())
	userId2, _ := NewUserId(uuid.NewString())
	email1, _ := NewEmail("test1@example.com")
	email2, _ := NewEmail("test2@example.com")
	password, _ := NewPassword("validPassword123")

	user1, _ := NewUser(userId1, "TestUser1", email1, password)
	user2, _ := NewUser(userId1, "TestUser2", email2, password)
	user3, _ := NewUser(userId2, "TestUser3", email1, password)

	if !user1.Equals(user2) {
		t.Errorf("Equals() should return true for users with same Id")
	}

	if user1.Equals(user3) {
		t.Errorf("Equals() should return false for users with different Ids")
	}
}

func TestUserString(t *testing.T) {
	userId, _ := NewUserId(uuid.NewString())
	name := "TestUser"
	email, _ := NewEmail("test@example.com")
	password, _ := NewPassword("validPassword123")

	user, _ := NewUser(userId, name, email, password)
	expected := "User{Id: " + userId.Value() + ", Name: " + name + ", Email: " + email.String() + "}"
	if user.String() != expected {
		t.Errorf("String() returned %s, expected %s", user.String(), expected)
	}
}