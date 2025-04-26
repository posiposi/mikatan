package model

import (
	"time"
)

type User struct {
	ID        string    `json:"id" gorm:"primary_key"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Readings  []Reading `gorm:"many2many:readings"`
}

type UserResponse struct {
	ID    string `json:"id" gorm:"primary_key"`
	Email string `json:"email" gorm:"unique"`
}
