package model

import (
	"time"
)

type Price struct {
	PriceId         string     `gorm:"column:price_id;primaryKey"`
	ItemId          string     `gorm:"column:item_id"`
	PriceWithTax    int        `gorm:"column:price_with_tax"`
	PriceWithoutTax int        `gorm:"column:price_without_tax"`
	TaxRate         float64    `gorm:"column:tax_rate"`
	Currency        string     `gorm:"column:currency"`
	StartDate       time.Time  `gorm:"column:start_date"`
	EndDate         *time.Time `gorm:"column:end_date"`
	CreatedAt       time.Time  `gorm:"column:created_at"`
	UpdatedAt       *time.Time `gorm:"column:updated_at"`
}
