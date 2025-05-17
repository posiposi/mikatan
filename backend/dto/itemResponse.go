package dto

import (
	"time"
)

type ItemResponse struct {
	ItemId      string    `json:"item_id"`
	UserId      string    `json:"user_id"`
	ItemName    string    `json:"item_name"`
	Stock       bool      `json:"stock"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
