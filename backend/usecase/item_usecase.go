// Package usecase provides business logic implementations.
package usecase

import (
	"github.com/posiposi/project/backend/domain"
	"github.com/posiposi/project/backend/dto"
	"github.com/posiposi/project/backend/model"
	"github.com/posiposi/project/backend/repository"
)

type IItemUsecase interface {
	GetAllItems() ([]dto.ItemResponse, error)
	CreateItem(item model.Item, userID string) (*dto.ItemResponse, error)
}

type itemUsecase struct {
	ir repository.IItemRepository
}

func NewItemUsecase(ir repository.IItemRepository) IItemUsecase {
	return &itemUsecase{ir}
}

func (iu *itemUsecase) GetAllItems() ([]dto.ItemResponse, error) {
	items, err := iu.ir.GetAllItems()
	if err != nil {
		return nil, err
	}
	itemsRes := &[]dto.ItemResponse{}
	for _, v := range items {
		t := v.ToDto()
		*itemsRes = append(*itemsRes, t)
	}
	return *itemsRes, nil
}

func (iu *itemUsecase) CreateItem(item model.Item, userID string) (*dto.ItemResponse, error) {
	userIdValue, err := domain.NewUserId(userID)
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
	domainItem, err := domain.NewItem(nil, *userIdValue, *itemName, *stock, *description)
	if err != nil {
		return nil, err
	}
	createdItem, err := iu.ir.CreateItem(domainItem)
	if err != nil {
		return nil, err
	}
	res := createdItem.ToDto()
	return &res, nil
}
