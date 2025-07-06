package model

type Item struct {
	ItemName    string `json:"item_name" validate:"required"`
	Stock       bool   `json:"stock"`
	Description string `json:"description"`
}