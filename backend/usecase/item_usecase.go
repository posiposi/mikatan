package usecase

import (
	"github.com/posiposi/project/backend/dto"
	"github.com/posiposi/project/backend/repository"
)

type IItemUsecase interface {
	GetAllItems() ([]dto.ItemResponse, error)
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
