package model

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	ItemId      string         `json:"itemId" gorm:"primaryKey"`
	UserId      string         `json:"userId" gorm:"size:36;not null"`
	ItemName    string         `json:"itemName" gorm:"not null"`
	Stock       bool           `json:"stock" gorm:"not null;default:true"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"createdAt" gorm:"not null"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	User        User
}
