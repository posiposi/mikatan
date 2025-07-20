// Package repository provides data persistence layer implementations.
package repository

import (
	"github.com/posiposi/project/backend/domain"
	"github.com/posiposi/project/backend/internal/orm/model"
	"gorm.io/gorm"
)

type IItemRepository interface {
	GetAllItems() (domain.Items, error)
	GetItemByID(itemId *domain.ItemId) (*domain.Item, error)
	CreateItem(item *domain.Item) (*domain.Item, error)
	UpdateItem(item *domain.Item) (*domain.Item, error)
	DeleteItem(itemId *domain.ItemId) error
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
		ItemId:      item.ItemId(),
		UserId:      item.UserId(),
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

func (ir *itemRepository) GetItemByID(itemId *domain.ItemId) (*domain.Item, error) {
	var ormItem model.Item
	if err := ir.db.Where("item_id = ?", itemId.Value()).First(&ormItem).Error; err != nil {
		return nil, err
	}

	itemIdValue, err := domain.NewItemId(ormItem.ItemId)
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

	item, err := domain.NewItem(itemIdValue, *userId, *itemName, *stock, *description)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (ir *itemRepository) UpdateItem(item *domain.Item) (*domain.Item, error) {
	ormItem := model.Item{
		ItemId:      item.ItemId(),
		UserId:      item.UserId(),
		ItemName:    item.ItemName(),
		Stock:       item.Stock(),
		Description: item.Description(),
	}

	result := ir.db.Where("item_id = ?", item.ItemId()).Select("item_name", "stock", "description").Updates(&ormItem)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	var updatedOrmItem model.Item
	if err := ir.db.Where("item_id = ?", item.ItemId()).First(&updatedOrmItem).Error; err != nil {
		return nil, err
	}

	itemId, err := domain.NewItemId(updatedOrmItem.ItemId)
	if err != nil {
		return nil, err
	}
	userId, err := domain.NewUserId(updatedOrmItem.UserId)
	if err != nil {
		return nil, err
	}
	itemName, err := domain.NewItemName(updatedOrmItem.ItemName)
	if err != nil {
		return nil, err
	}
	stock, err := domain.NewStock(updatedOrmItem.Stock)
	if err != nil {
		return nil, err
	}
	description, err := domain.NewDescription(updatedOrmItem.Description)
	if err != nil {
		return nil, err
	}

	updatedItem, err := domain.NewItem(itemId, *userId, *itemName, *stock, *description)
	if err != nil {
		return nil, err
	}

	return updatedItem, nil
}

func (ir *itemRepository) DeleteItem(itemId *domain.ItemId) error {
	result := ir.db.Where("item_id = ?", itemId.Value()).Delete(&model.Item{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
