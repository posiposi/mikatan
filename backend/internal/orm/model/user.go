package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserId    string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Items     []Item
}
