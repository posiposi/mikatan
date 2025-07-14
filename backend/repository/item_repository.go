// Package repository provides data persistence layer implementations.
package repository

import (
	"github.com/posiposi/project/backend/domain"
	"github.com/posiposi/project/backend/internal/orm/model"
	"gorm.io/gorm"
)

type IItemRepository interface {
	GetAllItems() (domain.Items, error)
	CreateItem(item *domain.Item) (*domain.Item, error)
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
		itemId, err := domain.NewItemId(v.ItemId)
		if err != nil {
			return nil, err
		}
		userId, err := domain.NewUserId(v.UserId)
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
			itemId,
			*userId,
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

func (ir *itemRepository) CreateItem(item *domain.Item) (*domain.Item, error) {
	ormItem := model.Item{
		ItemId:      item.ItemID(),
		UserId:      item.UserID(),
		ItemName:    item.ItemName(),
		Stock:       item.Stock(),
		Description: item.Description(),
	}

	if err := ir.db.Create(&ormItem).Error; err != nil {
		return nil, err
	}

	itemId, err := domain.NewItemId(ormItem.ItemId)
	if err != nil {
		return nil, err
	}
	userId, err := domain.NewUserId(ormItem.UserId)
	if err != nil {
		return nil, err
	}
	itemName, err := domain.NewItemName(ormItem.ItemName)
	if err != nil {
		return nil, err
	}
	stock, err := domain.NewStock(ormItem.Stock)
	if err != nil {
		return nil, err
	}
	description, err := domain.NewDescription(ormItem.Description)
	if err != nil {
		return nil, err
	}

	createdItem, err := domain.NewItem(itemId, *userId, *itemName, *stock, *description)
	if err != nil {
		return nil, err
	}

	return createdItem, nil
}
