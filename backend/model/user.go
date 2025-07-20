package model

import (
	"time"
)

type User struct {
	Id        string    `json:"id" gorm:"column:user_id;primaryKey"`
	Name      string    `json:"name" gorm:"column:name"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	Role      string    `json:"role" gorm:"default:USER"`
	IsAdmin   bool      `json:"is_admin" gorm:"default:false"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserResponse struct {
	Id    string `json:"id" gorm:"column:user_id;primaryKey"`
	Email string `json:"email" gorm:"unique"`
}
