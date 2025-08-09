package repository

import (
	"context"
	"errors"
	"time"

	"github.com/posiposi/project/backend/domain"
	"github.com/posiposi/project/backend/internal/orm/model"
	"gorm.io/gorm"
)

type PriceRepository interface {
	Create(ctx context.Context, price *domain.Price) error
	FindById(ctx context.Context, priceId string) (*domain.Price, error)
	FindByItemId(ctx context.Context, itemId string) ([]*domain.Price, error)
	FindCurrentByItemId(ctx context.Context, itemId string) (*domain.Price, error)
	UpdateByItemId(ctx context.Context, itemId string, price *domain.Price) error
	Delete(ctx context.Context, priceId string) error
}

type priceRepository struct {
	db *gorm.DB
}

func NewPriceRepository(db *gorm.DB) PriceRepository {
	return &priceRepository{db: db}
}

func (r *priceRepository) Create(ctx context.Context, price *domain.Price) error {
	updatedAt := price.UpdatedAt()
	dbPrice := &model.Price{
		PriceId:         price.PriceId(),
		ItemId:          price.ItemId(),
		PriceWithTax:    price.PriceWithTax(),
		PriceWithoutTax: price.PriceWithoutTax(),
		TaxRate:         price.TaxRate(),
		Currency:        price.Currency(),
		StartDate:       price.StartDate(),
		EndDate:         price.EndDate(),
		CreatedAt:       price.CreatedAt(),
		UpdatedAt:       &updatedAt,
	}

	if err := r.db.WithContext(ctx).Create(dbPrice).Error; err != nil {
		return err
	}

	return nil
}

func (r *priceRepository) FindById(ctx context.Context, priceId string) (*domain.Price, error) {
	var dbPrice model.Price
	if err := r.db.WithContext(ctx).Where("price_id = ?", priceId).First(&dbPrice).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	return r.toDomainPrice(&dbPrice)
}

func (r *priceRepository) FindByItemId(ctx context.Context, itemId string) ([]*domain.Price, error) {
	var dbPrices []model.Price
	if err := r.db.WithContext(ctx).Where("item_id = ?", itemId).Order("start_date DESC").Find(&dbPrices).Error; err != nil {
		return nil, err
	}

	prices := make([]*domain.Price, 0)
	for _, dbPrice := range dbPrices {
		price, err := r.toDomainPrice(&dbPrice)
		if err != nil {
			return nil, err
		}
		prices = append(prices, price)
	}

	return prices, nil
}

func (r *priceRepository) FindCurrentByItemId(ctx context.Context, itemId string) (*domain.Price, error) {
	var dbPrice model.Price
	now := time.Now()

	query := r.db.WithContext(ctx).
		Where("item_id = ?", itemId).
		Where("start_date <= ?", now).
		Where("end_date IS NULL OR end_date > ?", now).
		Order("start_date DESC")

	if err := query.First(&dbPrice).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return r.toDomainPrice(&dbPrice)
}

func (r *priceRepository) UpdateByItemId(ctx context.Context, itemId string, newPrice *domain.Price) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		now := time.Now()
		result := tx.Model(&model.Price{}).
			Where("item_id = ? AND end_date IS NULL", itemId).
			Update("end_date", now)

		if result.Error != nil {
			return result.Error
		}

		if result.RowsAffected == 0 {
			return errors.New("指定された商品の有効な料金が見つかりません")
		}

		updatedAt := newPrice.UpdatedAt()
		dbPrice := &model.Price{
			PriceId:         newPrice.PriceId(),
			ItemId:          newPrice.ItemId(),
			PriceWithTax:    newPrice.PriceWithTax(),
			PriceWithoutTax: newPrice.PriceWithoutTax(),
			TaxRate:         newPrice.TaxRate(),
			Currency:        newPrice.Currency(),
			StartDate:       newPrice.StartDate(),
			EndDate:         newPrice.EndDate(),
			CreatedAt:       newPrice.CreatedAt(),
			UpdatedAt:       &updatedAt,
		}

		if err := tx.Create(dbPrice).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *priceRepository) Delete(ctx context.Context, priceId string) error {
	result := r.db.WithContext(ctx).Where("price_id = ?", priceId).Delete(&model.Price{})
	return result.Error
}

func (r *priceRepository) toDomainPrice(dbPrice *model.Price) (*domain.Price, error) {
	priceId, err := domain.NewPriceId(dbPrice.PriceId)
	if err != nil {
		return nil, err
	}

	itemId, err := domain.NewItemId(dbPrice.ItemId)
	if err != nil {
		return nil, err
	}

	priceWithTax, err := domain.NewPriceWithTax(dbPrice.PriceWithTax)
	if err != nil {
		return nil, err
	}

	priceWithoutTax, err := domain.NewPriceWithoutTax(dbPrice.PriceWithoutTax)
	if err != nil {
		return nil, err
	}

	taxRate, err := domain.NewTaxRate(dbPrice.TaxRate)
	if err != nil {
		return nil, err
	}

	currency, err := domain.NewCurrency(dbPrice.Currency)
	if err != nil {
		return nil, err
	}

	return domain.NewPrice(
		priceId,
		*itemId,
		*priceWithTax,
		*priceWithoutTax,
		*taxRate,
		*currency,
		dbPrice.StartDate,
		dbPrice.EndDate,
	)
}
