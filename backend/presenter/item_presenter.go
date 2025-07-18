package presenter

import (
	"time"

	"github.com/posiposi/project/backend/domain"
)

type ItemResponseJSON struct {
	ItemId      string    `json:"item_id"`
	UserId      string    `json:"user_id"`
	ItemName    string    `json:"item_name"`
	Stock       bool      `json:"stock"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type IItemPresenter interface {
	ToJSON(item *domain.Item) ItemResponseJSON
	ToJSONList(items []*domain.Item) []ItemResponseJSON
}

type itemPresenter struct{}

func NewItemPresenter() IItemPresenter {
	return &itemPresenter{}
}

func (p *itemPresenter) ToJSON(item *domain.Item) ItemResponseJSON {
	return ItemResponseJSON{
		ItemId:      item.ItemId(),
		UserId:      item.UserId(),
		ItemName:    item.ItemName(),
		Stock:       item.Stock(),
		Description: item.Description(),
		CreatedAt:   item.CreatedAt(),
		UpdatedAt:   item.UpdatedAt(),
	}
}

func (p *itemPresenter) ToJSONList(items []*domain.Item) []ItemResponseJSON {
	result := make([]ItemResponseJSON, len(items))
	for i, item := range items {
		result[i] = p.ToJSON(item)
	}
	return result
}
