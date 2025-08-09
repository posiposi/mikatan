// Package usecase provides business logic implementations.
package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/posiposi/project/backend/domain"
	"github.com/posiposi/project/backend/repository"
	"github.com/posiposi/project/backend/usecase/request"
)

type PriceRepository interface {
	Create(ctx context.Context, price *domain.Price) error
	FindById(ctx context.Context, priceId string) (*domain.Price, error)
	FindByItemId(ctx context.Context, itemId string) ([]*domain.Price, error)
	FindCurrentByItemId(ctx context.Context, itemId string) (*domain.Price, error)
	UpdateByItemId(ctx context.Context, itemId string, price *domain.Price) error
	Delete(ctx context.Context, priceId string) error
}

type IItemUsecase interface {
	GetAllItems() ([]*domain.Item, error)
	GetItemByID(itemId string) (*domain.Item, error)
	CreateItem(req request.CreateItemRequest) (*domain.Item, error)
	UpdateItem(req request.UpdateItemRequest) (*domain.Item, error)
	DeleteItem(itemId string) error
}

type itemUsecase struct {
	ir repository.IItemRepository
	pr PriceRepository
}

func NewItemUsecase(ir repository.IItemRepository, pr PriceRepository) IItemUsecase {
	return &itemUsecase{ir, pr}
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

	// 料金設定がある場合は料金も作成
	if req.PriceWithTax != nil && req.PriceWithoutTax != nil && req.TaxRate != nil && req.Currency != nil {
		priceId, err := domain.NewPriceId(uuid.New().String())
		if err != nil {
			return nil, err
		}

		itemIdDomain, err := domain.NewItemId(createdItem.ItemId())
		if err != nil {
			return nil, err
		}

		priceWithTax, err := domain.NewPriceWithTax(*req.PriceWithTax)
		if err != nil {
			return nil, err
		}

		priceWithoutTax, err := domain.NewPriceWithoutTax(*req.PriceWithoutTax)
		if err != nil {
			return nil, err
		}

		taxRate, err := domain.NewTaxRate(*req.TaxRate)
		if err != nil {
			return nil, err
		}

		currency, err := domain.NewCurrency(*req.Currency)
		if err != nil {
			return nil, err
		}

		startDate := time.Now()
		price, err := domain.NewPrice(
			priceId,
			*itemIdDomain,
			*priceWithTax,
			*priceWithoutTax,
			*taxRate,
			*currency,
			startDate,
			nil,
		)
		if err != nil {
			return nil, err
		}

		err = iu.pr.Create(context.Background(), price)
		if err != nil {
			return nil, err
		}
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

	updatedItem, err := iu.ir.UpdateItem(updatedDomainItem)
	if err != nil {
		return nil, err
	}

	if req.PriceWithTax != nil && req.PriceWithoutTax != nil && req.TaxRate != nil && req.Currency != nil {
		priceId, err := domain.NewPriceId(uuid.New().String())
		if err != nil {
			return nil, err
		}

		priceWithTax, err := domain.NewPriceWithTax(*req.PriceWithTax)
		if err != nil {
			return nil, err
		}

		priceWithoutTax, err := domain.NewPriceWithoutTax(*req.PriceWithoutTax)
		if err != nil {
			return nil, err
		}

		taxRate, err := domain.NewTaxRate(*req.TaxRate)
		if err != nil {
			return nil, err
		}

		currency, err := domain.NewCurrency(*req.Currency)
		if err != nil {
			return nil, err
		}

		startDate := time.Now()
		price, err := domain.NewPrice(
			priceId,
			*itemId,
			*priceWithTax,
			*priceWithoutTax,
			*taxRate,
			*currency,
			startDate,
			nil,
		)
		if err != nil {
			return nil, err
		}

		err = iu.pr.UpdateByItemId(context.Background(), req.ItemId, price)
		if err != nil {
			return nil, err
		}
	}

	return updatedItem, nil
}

func (iu *itemUsecase) DeleteItem(itemId string) error {
	itemIdDomain, err := domain.NewItemId(itemId)
	if err != nil {
		return err
	}
	return iu.ir.DeleteItem(itemIdDomain)
}
