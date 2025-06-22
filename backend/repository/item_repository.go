// Package repository provides data persistence layer implementations.
package repository

import (
	"github.com/posiposi/project/backend/domain"
	"github.com/posiposi/project/backend/internal/orm/model"
	"gorm.io/gorm"
)

type IItemRepository interface {
	GetAllItems() (domain.Items, error)
}

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) IItemRepository {
	return &itemRepository{db}
}

func (ir *itemRepository) GetAllItems() (domain.Items, error) {
	// DBから全ての商品を取得（item_idで昇順ソート）
	var oi []model.Item
	if err := ir.db.Order("item_id ASC").Find(&oi).Error; err != nil {
		return nil, err
	}
	// ループ処理で各商品をドメインモデルに変換してから、itemsに追加する
	var items domain.Items
	for _, v := range oi {
		itemID, err := domain.NewItemID(v.ItemId)
		if err != nil {
			return nil, err
		}
		userID, err := domain.NewUserID(v.UserId)
		if err != nil {
			return nil, err
		}
		itemName, err := domain.NewItemName(v.ItemName)
		if err != nil {
			return nil, err
		}
		stock, err := domain.NewStock(v.Stock)
		if err != nil {
			return nil, err
		}
		description, err := domain.NewDescription(v.Description)
		if err != nil {
			return nil, err
		}
		item, err := domain.NewItem(
			itemID,
			*userID,
			*itemName,
			*stock,
			*description,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, *item)
	}
	return items, nil
}
