package domain

import (
	"fmt"
	"strings"
)

type User struct {
	id       *UserId
	name     string
	email    *Email
	password *Password
}

func NewUser(id *UserId, name string, email *Email, password *Password) (*User, error) {
	if id == nil {
		return nil, fmt.Errorf("user Id cannot be nil")
	}

	if strings.TrimSpace(name) == "" {
		return nil, fmt.Errorf("user name cannot be empty")
	}

	if email == nil {
		return nil, fmt.Errorf("email cannot be nil")
	}

	if password == nil {
		return nil, fmt.Errorf("password cannot be nil")
	}

	return &User{
		id:       id,
		name:     name,
		email:    email,
		password: password,
	}, nil
}

func (u *User) Id() *UserId {
	return u.id
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Email() *Email {
	return u.email
}

func (u *User) Password() *Password {
	return u.password
}

func (u *User) Equals(other *User) bool {
	if other == nil {
		return false
	}
	return u.id.Value() == other.id.Value()
}

func (u *User) String() string {
	return fmt.Sprintf("User{Id: %s, Name: %s, Email: %s}", u.id.Value(), u.name, u.email.String())
}