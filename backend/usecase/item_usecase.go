// Package usecase provides business logic implementations.
package usecase

import (
	"github.com/posiposi/project/backend/domain"
	"github.com/posiposi/project/backend/model"
	"github.com/posiposi/project/backend/repository"
)

type IItemUsecase interface {
	GetAllItems() ([]*domain.Item, error)
	CreateItem(item model.Item, userID string) (*domain.Item, error)
}

type itemUsecase struct {
	ir repository.IItemRepository
}

func NewItemUsecase(ir repository.IItemRepository) IItemUsecase {
	return &itemUsecase{ir}
}

func (iu *itemUsecase) GetAllItems() ([]*domain.Item, error) {
	items, err := iu.ir.GetAllItems()
	if err != nil {
		return nil, err
	}
	result := make([]*domain.Item, len(items))
	for i, item := range items {
		itemCopy := item
		result[i] = &itemCopy
	}
	return result, nil
}

func (iu *itemUsecase) CreateItem(item model.Item, userID string) (*domain.Item, error) {
	userIDValue, err := domain.NewUserID(userID)
	if err != nil {
		return nil, err
	}
	itemName, err := domain.NewItemName(item.ItemName)
	if err != nil {
		return nil, err
	}
	stock, err := domain.NewStock(item.Stock)
	if err != nil {
		return nil, err
	}
	description, err := domain.NewDescription(item.Description)
	if err != nil {
		return nil, err
	}
	domainItem, err := domain.NewItem(nil, *userIDValue, *itemName, *stock, *description)
	if err != nil {
		return nil, err
	}
	createdItem, err := iu.ir.CreateItem(domainItem)
	if err != nil {
		return nil, err
	}
	return createdItem, nil
}
