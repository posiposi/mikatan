// Package usecase provides business logic implementations.
package usecase

import (
	"github.com/posiposi/project/backend/domain"
	"github.com/posiposi/project/backend/repository"
	"github.com/posiposi/project/backend/usecase/request"
)

type IItemUsecase interface {
	GetAllItems() ([]*domain.Item, error)
	GetItemByID(itemId string) (*domain.Item, error)
	CreateItem(req request.CreateItemRequest) (*domain.Item, error)
	UpdateItem(req request.UpdateItemRequest) (*domain.Item, error)
	DeleteItem(itemId string) error
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

func (iu *itemUsecase) CreateItem(req request.CreateItemRequest) (*domain.Item, error) {
	userId, err := domain.NewUserId(req.UserId)
	if err != nil {
		return nil, err
	}
	
	itemName, err := domain.NewItemName(req.ItemName)
	if err != nil {
		return nil, err
	}
	
	stock, err := domain.NewStock(req.Stock)
	if err != nil {
		return nil, err
	}
	
	description, err := domain.NewDescription(req.Description)
	if err != nil {
		return nil, err
	}
	
	domainItem, err := domain.NewItem(nil, *userId, *itemName, *stock, *description)
	if err != nil {
		return nil, err
	}
	
	createdItem, err := iu.ir.CreateItem(domainItem)
	if err != nil {
		return nil, err
	}
	return createdItem, nil
}

func (iu *itemUsecase) GetItemByID(itemId string) (*domain.Item, error) {
	itemIdDomain, err := domain.NewItemId(itemId)
	if err != nil {
		return nil, err
	}
	return iu.ir.GetItemByID(itemIdDomain)
}

func (iu *itemUsecase) UpdateItem(req request.UpdateItemRequest) (*domain.Item, error) {
	itemId, err := domain.NewItemId(req.ItemId)
	if err != nil {
		return nil, err
	}
	
	itemName, err := domain.NewItemName(req.ItemName)
	if err != nil {
		return nil, err
	}
	
	stock, err := domain.NewStock(req.Stock)
	if err != nil {
		return nil, err
	}
	
	description, err := domain.NewDescription(req.Description)
	if err != nil {
		return nil, err
	}
	
	existingItem, err := iu.ir.GetItemByID(itemId)
	if err != nil {
		return nil, err
	}
	
	userId, err := domain.NewUserId(existingItem.UserId())
	if err != nil {
		return nil, err
	}
	
	updatedDomainItem, err := domain.NewItem(itemId, *userId, *itemName, *stock, *description)
	if err != nil {
		return nil, err
	}
	
	return iu.ir.UpdateItem(updatedDomainItem)
}

func (iu *itemUsecase) DeleteItem(itemId string) error {
	itemIdDomain, err := domain.NewItemId(itemId)
	if err != nil {
		return err
	}
	return iu.ir.DeleteItem(itemIdDomain)
}
