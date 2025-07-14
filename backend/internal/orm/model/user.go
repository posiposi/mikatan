package model

import (
	"time"
)

type User struct {
	UserId    string    `json:"user_id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Email     string    `json:"email" gorm:"type:varchar(255);not null;unique"`
	Password  string    `json:"password" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at"`
	Items     []Item
}
