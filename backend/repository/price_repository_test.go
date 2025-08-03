package repository

import (
	"context"
	"testing"
	"time"

	"github.com/posiposi/project/backend/testutil"

	"github.com/posiposi/project/backend/domain"

	"github.com/stretchr/testify/assert"
)

func TestPriceRepository_Create(t *testing.T) {
	db := testutil.NewTestDB(t)
	repo := NewPriceRepository(db)

	t.Run("正常に料金を作成できる", func(t *testing.T) {
		priceId, _ := domain.NewPriceId("550e8400-e29b-41d4-a716-446655440000")
		itemId, _ := domain.NewItemId("550e8400-e29b-41d4-a716-446655440001")
		priceWithTax, _ := domain.NewPriceWithTax(1100)
		priceWithoutTax, _ := domain.NewPriceWithoutTax(1000)
		taxRate, _ := domain.NewTaxRate(10.0)
		currency, _ := domain.NewCurrency("JPY")
		startDate := time.Now()

		price, _ := domain.NewPrice(
			priceId,
			*itemId,
			*priceWithTax,
			*priceWithoutTax,
			*taxRate,
			*currency,
			startDate,
			nil,
		)

		err := repo.Create(context.Background(), price)

		assert.NoError(t, err)

		found, err := repo.FindById(context.Background(), price.PriceId())
		assert.NoError(t, err)
		assert.NotNil(t, found)
		assert.Equal(t, price.PriceId(), found.PriceId())
	})

	t.Run("同じIDの料金は作成できない", func(t *testing.T) {
		priceId, _ := domain.NewPriceId("550e8400-e29b-41d4-a716-446655440002")
		itemId, _ := domain.NewItemId("550e8400-e29b-41d4-a716-446655440001")
		priceWithTax, _ := domain.NewPriceWithTax(1100)
		priceWithoutTax, _ := domain.NewPriceWithoutTax(1000)
		taxRate, _ := domain.NewTaxRate(10.0)
		currency, _ := domain.NewCurrency("JPY")
		startDate := time.Now()

		price, _ := domain.NewPrice(
			priceId,
			*itemId,
			*priceWithTax,
			*priceWithoutTax,
			*taxRate,
			*currency,
			startDate,
			nil,
		)

		err := repo.Create(context.Background(), price)
		assert.NoError(t, err)

		err = repo.Create(context.Background(), price)

		assert.Error(t, err)
	})
}

func TestPriceRepository_FindById(t *testing.T) {
	db := testutil.NewTestDB(t)
	repo := NewPriceRepository(db)

	t.Run("存在する料金を取得できる", func(t *testing.T) {
		priceId, _ := domain.NewPriceId("550e8400-e29b-41d4-a716-446655440003")
		itemId, _ := domain.NewItemId("550e8400-e29b-41d4-a716-446655440001")
		priceWithTax, _ := domain.NewPriceWithTax(2200)
		priceWithoutTax, _ := domain.NewPriceWithoutTax(2000)
		taxRate, _ := domain.NewTaxRate(10.0)
		currency, _ := domain.NewCurrency("JPY")
		startDate := time.Now()

		price, _ := domain.NewPrice(
			priceId,
			*itemId,
			*priceWithTax,
			*priceWithoutTax,
			*taxRate,
			*currency,
			startDate,
			nil,
		)
		repo.Create(context.Background(), price)

		found, err := repo.FindById(context.Background(), price.PriceId())

		assert.NoError(t, err)
		assert.NotNil(t, found)
		assert.Equal(t, price.PriceId(), found.PriceId())
		assert.Equal(t, price.ItemId(), found.ItemId())
		assert.Equal(t, price.PriceWithTax(), found.PriceWithTax())
		assert.Equal(t, price.PriceWithoutTax(), found.PriceWithoutTax())
	})

	t.Run("存在しない料金の場合、nilが返される", func(t *testing.T) {
		found, err := repo.FindById(context.Background(), "non-existent-id")

		assert.NoError(t, err)
		assert.Nil(t, found)
	})
}

func TestPriceRepository_FindByItemId(t *testing.T) {
	db := testutil.NewTestDB(t)
	repo := NewPriceRepository(db)

	t.Run("特定の商品の料金一覧を取得できる", func(t *testing.T) {
		itemId, _ := domain.NewItemId("550e8400-e29b-41d4-a716-446655440001")

		priceId1, _ := domain.NewPriceId("550e8400-e29b-41d4-a716-446655440004")
		priceWithTax1, _ := domain.NewPriceWithTax(1100)
		priceWithoutTax1, _ := domain.NewPriceWithoutTax(1000)
		taxRate, _ := domain.NewTaxRate(10.0)
		currency, _ := domain.NewCurrency("JPY")
		startDate1 := time.Now().AddDate(0, -1, 0)

		price1, _ := domain.NewPrice(
			priceId1,
			*itemId,
			*priceWithTax1,
			*priceWithoutTax1,
			*taxRate,
			*currency,
			startDate1,
			nil,
		)
		priceId2, _ := domain.NewPriceId("550e8400-e29b-41d4-a716-446655440005")
		priceWithTax2, _ := domain.NewPriceWithTax(2200)
		priceWithoutTax2, _ := domain.NewPriceWithoutTax(2000)
		startDate2 := time.Now()
		price2, _ := domain.NewPrice(
			priceId2,
			*itemId,
			*priceWithTax2,
			*priceWithoutTax2,
			*taxRate,
			*currency,
			startDate2,
			nil,
		)

		repo.Create(context.Background(), price1)
		repo.Create(context.Background(), price2)

		// Act
		prices, err := repo.FindByItemId(context.Background(), itemId.Value())

		// Assert
		assert.NoError(t, err)
		assert.Len(t, prices, 2)
	})

	t.Run("料金が存在しない商品の場合、空配列が返される", func(t *testing.T) {
		prices, err := repo.FindByItemId(context.Background(), "non-existent-item-id")

		assert.NoError(t, err)
		assert.Empty(t, prices)
	})
}

func TestPriceRepository_FindCurrentByItemId(t *testing.T) {
	db := testutil.NewTestDB(t)
	repo := NewPriceRepository(db)

	t.Run("現在有効な料金を取得できる", func(t *testing.T) {
		itemId, _ := domain.NewItemId("550e8400-e29b-41d4-a716-446655440001")
		taxRate, _ := domain.NewTaxRate(10.0)
		currency, _ := domain.NewCurrency("JPY")

		priceId1, _ := domain.NewPriceId("550e8400-e29b-41d4-a716-446655440006")
		priceWithTax1, _ := domain.NewPriceWithTax(1100)
		priceWithoutTax1, _ := domain.NewPriceWithoutTax(1000)
		startDate1 := time.Now().AddDate(0, -2, 0)
		endDate1 := time.Now().AddDate(0, -1, 0)

		price1, _ := domain.NewPrice(
			priceId1,
			*itemId,
			*priceWithTax1,
			*priceWithoutTax1,
			*taxRate,
			*currency,
			startDate1,
			&endDate1,
		)
		priceId2, _ := domain.NewPriceId("550e8400-e29b-41d4-a716-446655440007")
		priceWithTax2, _ := domain.NewPriceWithTax(2200)
		priceWithoutTax2, _ := domain.NewPriceWithoutTax(2000)
		startDate2 := time.Now().AddDate(0, -1, 0)
		price2, _ := domain.NewPrice(
			priceId2,
			*itemId,
			*priceWithTax2,
			*priceWithoutTax2,
			*taxRate,
			*currency,
			startDate2,
			nil,
		)

		repo.Create(context.Background(), price1)
		repo.Create(context.Background(), price2)

		current, err := repo.FindCurrentByItemId(context.Background(), itemId.Value())

		assert.NoError(t, err)
		assert.NotNil(t, current)
		assert.Equal(t, price2.PriceId(), current.PriceId())
		assert.Equal(t, 2200, current.PriceWithTax())
	})

	t.Run("有効な料金が存在しない場合、nilが返される", func(t *testing.T) {
		itemId, _ := domain.NewItemId("550e8400-e29b-41d4-a716-446655440002")
		priceId, _ := domain.NewPriceId("550e8400-e29b-41d4-a716-446655440008")
		priceWithTax, _ := domain.NewPriceWithTax(1100)
		priceWithoutTax, _ := domain.NewPriceWithoutTax(1000)
		taxRate, _ := domain.NewTaxRate(10.0)
		currency, _ := domain.NewCurrency("JPY")
		startDate := time.Now().AddDate(0, -2, 0)
		endDate := time.Now().AddDate(0, -1, 0)

		price, _ := domain.NewPrice(
			priceId,
			*itemId,
			*priceWithTax,
			*priceWithoutTax,
			*taxRate,
			*currency,
			startDate,
			&endDate,
		)

		repo.Create(context.Background(), price)

		current, err := repo.FindCurrentByItemId(context.Background(), itemId.Value())

		assert.NoError(t, err)
		assert.Nil(t, current)
	})
}

func TestPriceRepository_Update(t *testing.T) {
	db := testutil.NewTestDB(t)
	repo := NewPriceRepository(db)

	t.Run("料金を更新できる", func(t *testing.T) {
		priceId, _ := domain.NewPriceId("550e8400-e29b-41d4-a716-446655440009")
		itemId, _ := domain.NewItemId("550e8400-e29b-41d4-a716-446655440001")
		priceWithTax, _ := domain.NewPriceWithTax(1100)
		priceWithoutTax, _ := domain.NewPriceWithoutTax(1000)
		taxRate, _ := domain.NewTaxRate(10.0)
		currency, _ := domain.NewCurrency("JPY")
		startDate := time.Now()

		price, _ := domain.NewPrice(
			priceId,
			*itemId,
			*priceWithTax,
			*priceWithoutTax,
			*taxRate,
			*currency,
			startDate,
			nil,
		)
		repo.Create(context.Background(), price)

		newPriceWithTax, _ := domain.NewPriceWithTax(2200)
		newPriceWithoutTax, _ := domain.NewPriceWithoutTax(2000)
		updatedPrice, _ := domain.NewPrice(
			priceId,
			*itemId,
			*newPriceWithTax,
			*newPriceWithoutTax,
			*taxRate,
			*currency,
			startDate,
			nil,
		)

		err := repo.Update(context.Background(), updatedPrice)

		assert.NoError(t, err)

		found, _ := repo.FindById(context.Background(), price.PriceId())
		assert.Equal(t, 2200, found.PriceWithTax())
		assert.Equal(t, 2000, found.PriceWithoutTax())
	})

	t.Run("存在しない料金の更新はエラーになる", func(t *testing.T) {
		priceId, _ := domain.NewPriceId("non-existent-price-id")
		itemId, _ := domain.NewItemId("550e8400-e29b-41d4-a716-446655440001")
		priceWithTax, _ := domain.NewPriceWithTax(1100)
		priceWithoutTax, _ := domain.NewPriceWithoutTax(1000)
		taxRate, _ := domain.NewTaxRate(10.0)
		currency, _ := domain.NewCurrency("JPY")

		price, _ := domain.NewPrice(
			priceId,
			*itemId,
			*priceWithTax,
			*priceWithoutTax,
			*taxRate,
			*currency,
			time.Now(),
			nil,
		)

		err := repo.Update(context.Background(), price)

		assert.Error(t, err)
	})
}

func TestPriceRepository_Delete(t *testing.T) {
	db := testutil.NewTestDB(t)
	repo := NewPriceRepository(db)

	t.Run("料金を削除できる", func(t *testing.T) {
		priceId, _ := domain.NewPriceId("550e8400-e29b-41d4-a716-446655440010")
		itemId, _ := domain.NewItemId("550e8400-e29b-41d4-a716-446655440001")
		priceWithTax, _ := domain.NewPriceWithTax(1100)
		priceWithoutTax, _ := domain.NewPriceWithoutTax(1000)
		taxRate, _ := domain.NewTaxRate(10.0)
		currency, _ := domain.NewCurrency("JPY")

		price, _ := domain.NewPrice(
			priceId,
			*itemId,
			*priceWithTax,
			*priceWithoutTax,
			*taxRate,
			*currency,
			time.Now(),
			nil,
		)
		repo.Create(context.Background(), price)

		err := repo.Delete(context.Background(), price.PriceId())

		assert.NoError(t, err)

		found, _ := repo.FindById(context.Background(), price.PriceId())
		assert.Nil(t, found)
	})

	t.Run("存在しない料金の削除でもエラーにならない", func(t *testing.T) {
		err := repo.Delete(context.Background(), "non-existent-id")

		assert.NoError(t, err)
	})
}
