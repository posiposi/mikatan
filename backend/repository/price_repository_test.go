package repository

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/posiposi/project/backend/domain"
	"github.com/posiposi/project/backend/internal/orm/model"
	"github.com/posiposi/project/backend/testutil"
	"github.com/stretchr/testify/assert"
)

const (
	testUserId = "f47ac10b-58cc-4372-a567-0e02b2c3d111"
	testItemId = "f47ac10b-58cc-4372-a567-0e02b2c3d110"
)

func setupPriceTestDB() *gorm.DB {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalln("Error loading .env file")
	}
	mysqlPassword := os.Getenv("MYSQL_TEST_ROOT_PASSWORD")
	mysqlHost := os.Getenv("MYSQL_TEST_HOST")
	mysqlDatabase := os.Getenv("MYSQL_TEST_DATABASE")
	dsn := fmt.Sprintf("root:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlPassword,
		mysqlHost,
		mysqlDatabase,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to connect to test database:", err)
	}

	err = testutil.RunPrismaMigrationForTest()
	if err != nil {
		log.Fatalln("Failed to run Prisma migration:", err)
	}

	return db
}

func seedPriceTestData(db *gorm.DB) error {
	user := model.User{Id: testUserId, Name: "User1", Email: "user1@example.com", Password: "password", Role: "USER", IsAdmin: false}
	item := model.Item{ItemId: testItemId, UserId: testUserId, ItemName: "Item1", Stock: true, Description: "Description1"}

	if err := db.Create(&user).Error; err != nil {
		return err
	}

	if err := db.Create(&item).Error; err != nil {
		return err
	}

	return nil
}

func createTestModelPrice(priceWithTax, priceWithoutTax int, startDate time.Time, endDate *time.Time) model.Price {
	now := time.Now()
	return model.Price{
		PriceId:         uuid.New().String(),
		ItemId:          testItemId,
		PriceWithTax:    priceWithTax,
		PriceWithoutTax: priceWithoutTax,
		TaxRate:         10.0,
		Currency:        "JPY",
		StartDate:       startDate,
		EndDate:         endDate,
		CreatedAt:       now,
		UpdatedAt:       &now,
	}
}

func createTestDomainPrice(priceWithTax, priceWithoutTax int, startDate time.Time, endDate *time.Time) (*domain.Price, error) {
	priceId, err := domain.NewPriceId(uuid.New().String())
	if err != nil {
		return nil, err
	}

	itemId, err := domain.NewItemId(testItemId)
	if err != nil {
		return nil, err
	}

	priceWithTaxValue, err := domain.NewPriceWithTax(priceWithTax)
	if err != nil {
		return nil, err
	}

	priceWithoutTaxValue, err := domain.NewPriceWithoutTax(priceWithoutTax)
	if err != nil {
		return nil, err
	}

	taxRate, err := domain.NewTaxRate(10.0)
	if err != nil {
		return nil, err
	}

	currency, err := domain.NewCurrency("JPY")
	if err != nil {
		return nil, err
	}

	return domain.NewPrice(
		priceId,
		*itemId,
		*priceWithTaxValue,
		*priceWithoutTaxValue,
		*taxRate,
		*currency,
		startDate,
		endDate,
	)
}

func TestPriceRepository_Create(t *testing.T) {
	db := setupPriceTestDB()

	t.Run("Create Price - Success", func(t *testing.T) {
		tx := db.Begin()
		defer tx.Rollback()

		err := seedPriceTestData(tx)
		assert.NoError(t, err)

		price, err := createTestDomainPrice(1100, 1000, time.Now().AddDate(0, 0, -1), nil)
		assert.NoError(t, err)

		repo := NewPriceRepository(tx)
		err = repo.Create(context.Background(), price)
		assert.NoError(t, err)

		var savedPrice model.Price
		err = tx.Where("price_id = ?", price.PriceId()).First(&savedPrice).Error
		assert.NoError(t, err)
		assert.Equal(t, price.PriceId(), savedPrice.PriceId)
		assert.Equal(t, price.ItemId(), savedPrice.ItemId)
		assert.Equal(t, price.PriceWithTax(), savedPrice.PriceWithTax)
		assert.Equal(t, price.PriceWithoutTax(), savedPrice.PriceWithoutTax)
		assert.Equal(t, price.TaxRate(), savedPrice.TaxRate)
		assert.Equal(t, price.Currency(), savedPrice.Currency)
		assert.Equal(t, price.StartDate().Unix(), savedPrice.StartDate.Unix())
		assert.Nil(t, savedPrice.EndDate)
	})

	t.Run("Create Price - With EndDate", func(t *testing.T) {
		tx := db.Begin()
		defer tx.Rollback()

		err := seedPriceTestData(tx)
		assert.NoError(t, err)

		endDate := time.Now().AddDate(0, 0, 1)
		price, err := createTestDomainPrice(2200, 2000, time.Now().AddDate(0, 0, -2), &endDate)
		assert.NoError(t, err)

		repo := NewPriceRepository(tx)
		err = repo.Create(context.Background(), price)
		assert.NoError(t, err)

		var savedPrice model.Price
		err = tx.Where("price_id = ?", price.PriceId()).First(&savedPrice).Error
		assert.NoError(t, err)
		assert.NotNil(t, savedPrice.EndDate)
		assert.Equal(t, endDate.Unix(), savedPrice.EndDate.Unix())
	})
}

func TestPriceRepository_FindById(t *testing.T) {
	db := setupPriceTestDB()

	t.Run("FindById - Success", func(t *testing.T) {
		tx := db.Begin()
		defer tx.Rollback()

		err := seedPriceTestData(tx)
		assert.NoError(t, err)

		testPrice := createTestModelPrice(1100, 1000, time.Now().AddDate(0, 0, -1), nil)
		err = tx.Create(&testPrice).Error
		assert.NoError(t, err)

		repo := NewPriceRepository(tx)
		result, err := repo.FindById(context.Background(), testPrice.PriceId)
		assert.NoError(t, err)
		assert.NotNil(t, result)

		assert.Equal(t, testPrice.PriceId, result.PriceId())
		assert.Equal(t, testItemId, result.ItemId())
		assert.Equal(t, 1100, result.PriceWithTax())
		assert.Equal(t, 1000, result.PriceWithoutTax())
		assert.Equal(t, 10.0, result.TaxRate())
		assert.Equal(t, "JPY", result.Currency())
		assert.Nil(t, result.EndDate())
	})

	t.Run("FindById - With EndDate", func(t *testing.T) {
		tx := db.Begin()
		defer tx.Rollback()

		err := seedPriceTestData(tx)
		assert.NoError(t, err)

		endDate := time.Now().AddDate(0, 0, 1)
		testPrice := createTestModelPrice(2200, 2000, time.Now().AddDate(0, 0, -2), &endDate)
		err = tx.Create(&testPrice).Error
		assert.NoError(t, err)

		repo := NewPriceRepository(tx)
		result, err := repo.FindById(context.Background(), testPrice.PriceId)
		assert.NoError(t, err)
		assert.NotNil(t, result)

		assert.NotNil(t, result.EndDate())
		assert.Equal(t, endDate.Unix(), result.EndDate().Unix())
	})

	t.Run("FindById - Not Found", func(t *testing.T) {
		tx := db.Begin()
		defer tx.Rollback()

		repo := NewPriceRepository(tx)
		nonExistentId := uuid.New().String()

		result, err := repo.FindById(context.Background(), nonExistentId)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, gorm.ErrRecordNotFound, err)
	})
}

func TestPriceRepository_FindByItemId(t *testing.T) {
	db := setupPriceTestDB()

	t.Run("FindByItemId - Multiple Prices", func(t *testing.T) {
		tx := db.Begin()
		defer tx.Rollback()

		err := seedPriceTestData(tx)
		assert.NoError(t, err)

		prices := []model.Price{
			createTestModelPrice(1100, 1000, time.Now().AddDate(0, -2, 0), nil),
			createTestModelPrice(2200, 2000, time.Now().AddDate(0, -1, 0), nil),
			createTestModelPrice(3300, 3000, time.Now(), nil),
		}

		for _, price := range prices {
			err = tx.Create(&price).Error
			assert.NoError(t, err)
		}

		repo := NewPriceRepository(tx)
		results, err := repo.FindByItemId(context.Background(), testItemId)
		assert.NoError(t, err)
		assert.NotNil(t, results)
		assert.Len(t, results, 3)

		assert.Equal(t, 3300, results[0].PriceWithTax())
		assert.Equal(t, 2200, results[1].PriceWithTax())
		assert.Equal(t, 1100, results[2].PriceWithTax())
	})

	t.Run("FindByItemId - No Prices", func(t *testing.T) {
		tx := db.Begin()
		defer tx.Rollback()

		err := seedPriceTestData(tx)
		assert.NoError(t, err)

		repo := NewPriceRepository(tx)
		results, err := repo.FindByItemId(context.Background(), "non-existent-item-id")
		assert.NoError(t, err)
		assert.NotNil(t, results)
		assert.Len(t, results, 0)
	})

	t.Run("FindByItemId - With EndDate", func(t *testing.T) {
		tx := db.Begin()
		defer tx.Rollback()

		err := seedPriceTestData(tx)
		assert.NoError(t, err)

		endDate1 := time.Now().AddDate(0, -1, 0)
		endDate2 := time.Now().AddDate(0, 0, 15)

		prices := []model.Price{
			createTestModelPrice(1100, 1000, time.Now().AddDate(0, -3, 0), &endDate1),
			createTestModelPrice(2200, 2000, time.Now().AddDate(0, -1, 0), &endDate2),
		}

		for _, price := range prices {
			err = tx.Create(&price).Error
			assert.NoError(t, err)
		}

		repo := NewPriceRepository(tx)
		results, err := repo.FindByItemId(context.Background(), testItemId)
		assert.NoError(t, err)
		assert.Len(t, results, 2)

		assert.NotNil(t, results[0].EndDate())
		assert.NotNil(t, results[1].EndDate())
	})
}

func TestPriceRepository_FindCurrentByItemId(t *testing.T) {
	db := setupPriceTestDB()

	t.Run("FindCurrentByItemId - Current Price", func(t *testing.T) {
		tx := db.Begin()
		defer tx.Rollback()

		err := seedPriceTestData(tx)
		assert.NoError(t, err)

		pastEndDate := time.Now().AddDate(0, -1, 0)
		pastPrice := createTestModelPrice(1100, 1000, time.Now().AddDate(0, -3, 0), &pastEndDate)
		currentPrice := createTestModelPrice(2200, 2000, time.Now().AddDate(0, -1, 0), nil)
		futurePrice := createTestModelPrice(3300, 3000, time.Now().AddDate(0, 0, 1), nil)

		err = tx.Create(&pastPrice).Error
		assert.NoError(t, err)
		err = tx.Create(&currentPrice).Error
		assert.NoError(t, err)
		err = tx.Create(&futurePrice).Error
		assert.NoError(t, err)

		repo := NewPriceRepository(tx)
		result, err := repo.FindCurrentByItemId(context.Background(), testItemId)
		assert.NoError(t, err)
		assert.NotNil(t, result)

		assert.Equal(t, 2200, result.PriceWithTax())
		assert.Equal(t, 2000, result.PriceWithoutTax())
		assert.Nil(t, result.EndDate())
	})

	t.Run("FindCurrentByItemId - With Future EndDate", func(t *testing.T) {
		tx := db.Begin()
		defer tx.Rollback()

		err := seedPriceTestData(tx)
		assert.NoError(t, err)

		futureEndDate := time.Now().AddDate(0, 1, 0)
		currentPrice := createTestModelPrice(2200, 2000, time.Now().AddDate(0, -1, 0), &futureEndDate)

		err = tx.Create(&currentPrice).Error
		assert.NoError(t, err)

		repo := NewPriceRepository(tx)
		result, err := repo.FindCurrentByItemId(context.Background(), testItemId)
		assert.NoError(t, err)
		assert.NotNil(t, result)

		assert.Equal(t, 2200, result.PriceWithTax())
		assert.NotNil(t, result.EndDate())
	})

	t.Run("FindCurrentByItemId - No Current Price", func(t *testing.T) {
		tx := db.Begin()
		defer tx.Rollback()

		err := seedPriceTestData(tx)
		assert.NoError(t, err)

		pastEndDate := time.Now().AddDate(0, 0, -1)
		pastPrice := createTestModelPrice(1100, 1000, time.Now().AddDate(0, -2, 0), &pastEndDate)

		err = tx.Create(&pastPrice).Error
		assert.NoError(t, err)

		repo := NewPriceRepository(tx)
		result, err := repo.FindCurrentByItemId(context.Background(), testItemId)
		assert.NoError(t, err)
		assert.Nil(t, result)
	})

	t.Run("FindCurrentByItemId - No Prices", func(t *testing.T) {
		tx := db.Begin()
		defer tx.Rollback()

		err := seedPriceTestData(tx)
		assert.NoError(t, err)

		repo := NewPriceRepository(tx)
		result, err := repo.FindCurrentByItemId(context.Background(), "non-existent-item-id")
		assert.NoError(t, err)
		assert.Nil(t, result)
	})
}

func TestPriceRepository_UpdateByItemId(t *testing.T) {
	db := setupPriceTestDB()

	t.Run("UpdateByItemId - Success", func(t *testing.T) {
		tx := db.Begin()
		defer tx.Rollback()

		err := seedPriceTestData(tx)
		assert.NoError(t, err)

		oldPrice := createTestModelPrice(1100, 1000, time.Now().AddDate(0, -1, 0), nil)
		err = tx.Create(&oldPrice).Error
		assert.NoError(t, err)

		newPrice, err := createTestDomainPrice(2200, 2000, time.Now(), nil)
		assert.NoError(t, err)

		repo := NewPriceRepository(tx)
		err = repo.UpdateByItemId(context.Background(), testItemId, newPrice)
		assert.NoError(t, err)

		var oldPriceUpdated model.Price
		err = tx.Where("price_id = ?", oldPrice.PriceId).First(&oldPriceUpdated).Error
		assert.NoError(t, err)
		assert.NotNil(t, oldPriceUpdated.EndDate)

		var newPriceCreated model.Price
		err = tx.Where("price_id = ?", newPrice.PriceId()).First(&newPriceCreated).Error
		assert.NoError(t, err)
		assert.Equal(t, 2200, newPriceCreated.PriceWithTax)
		assert.Equal(t, 2000, newPriceCreated.PriceWithoutTax)
		assert.Nil(t, newPriceCreated.EndDate)
	})

	t.Run("UpdateByItemId - Multiple Active Prices", func(t *testing.T) {
		tx := db.Begin()
		defer tx.Rollback()

		err := seedPriceTestData(tx)
		assert.NoError(t, err)

		price1 := createTestModelPrice(1100, 1000, time.Now().AddDate(0, -2, 0), nil)
		price2 := createTestModelPrice(2200, 2000, time.Now().AddDate(0, -1, 0), nil)
		err = tx.Create(&price1).Error
		assert.NoError(t, err)
		err = tx.Create(&price2).Error
		assert.NoError(t, err)

		newPrice, err := createTestDomainPrice(3300, 3000, time.Now(), nil)
		assert.NoError(t, err)

		repo := NewPriceRepository(tx)
		err = repo.UpdateByItemId(context.Background(), testItemId, newPrice)
		assert.NoError(t, err)

		var count int64
		tx.Model(&model.Price{}).Where("item_id = ? AND end_date IS NOT NULL", testItemId).Count(&count)
		assert.Equal(t, int64(2), count)

		var activeCount int64
		tx.Model(&model.Price{}).Where("item_id = ? AND end_date IS NULL", testItemId).Count(&activeCount)
		assert.Equal(t, int64(1), activeCount)
	})

	t.Run("UpdateByItemId - No Active Price", func(t *testing.T) {
		tx := db.Begin()
		defer tx.Rollback()

		err := seedPriceTestData(tx)
		assert.NoError(t, err)

		pastEndDate := time.Now().AddDate(0, 0, -1)
		oldPrice := createTestModelPrice(1100, 1000, time.Now().AddDate(0, -2, 0), &pastEndDate)
		err = tx.Create(&oldPrice).Error
		assert.NoError(t, err)

		newPrice, err := createTestDomainPrice(2200, 2000, time.Now(), nil)
		assert.NoError(t, err)

		repo := NewPriceRepository(tx)
		err = repo.UpdateByItemId(context.Background(), testItemId, newPrice)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "指定された商品の有効な料金が見つかりません")
	})

	t.Run("UpdateByItemId - No Prices", func(t *testing.T) {
		tx := db.Begin()
		defer tx.Rollback()

		err := seedPriceTestData(tx)
		assert.NoError(t, err)

		newPrice, err := createTestDomainPrice(2200, 2000, time.Now(), nil)
		assert.NoError(t, err)

		repo := NewPriceRepository(tx)
		err = repo.UpdateByItemId(context.Background(), testItemId, newPrice)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "指定された商品の有効な料金が見つかりません")
	})
}

func TestPriceRepository_Delete(t *testing.T) {
	db := setupPriceTestDB()

	t.Run("Delete - Success", func(t *testing.T) {
		tx := db.Begin()
		defer tx.Rollback()

		err := seedPriceTestData(tx)
		assert.NoError(t, err)

		price := createTestModelPrice(1100, 1000, time.Now().AddDate(0, -1, 0), nil)
		err = tx.Create(&price).Error
		assert.NoError(t, err)

		repo := NewPriceRepository(tx)
		err = repo.Delete(context.Background(), price.PriceId)
		assert.NoError(t, err)

		var count int64
		err = tx.Model(&model.Price{}).Where("price_id = ?", price.PriceId).Count(&count).Error
		assert.NoError(t, err)
		assert.Equal(t, int64(0), count)
	})

	t.Run("Delete - Non Existent Price", func(t *testing.T) {
		tx := db.Begin()
		defer tx.Rollback()

		err := seedPriceTestData(tx)
		assert.NoError(t, err)

		repo := NewPriceRepository(tx)
		nonExistentId := uuid.New().String()
		err = repo.Delete(context.Background(), nonExistentId)
		assert.NoError(t, err)
	})

	t.Run("Delete - Multiple Prices", func(t *testing.T) {
		tx := db.Begin()
		defer tx.Rollback()

		err := seedPriceTestData(tx)
		assert.NoError(t, err)

		price1 := createTestModelPrice(1100, 1000, time.Now().AddDate(0, -2, 0), nil)
		price2 := createTestModelPrice(2200, 2000, time.Now().AddDate(0, -1, 0), nil)
		price3 := createTestModelPrice(3300, 3000, time.Now(), nil)

		err = tx.Create(&price1).Error
		assert.NoError(t, err)
		err = tx.Create(&price2).Error
		assert.NoError(t, err)
		err = tx.Create(&price3).Error
		assert.NoError(t, err)

		repo := NewPriceRepository(tx)
		err = repo.Delete(context.Background(), price2.PriceId)
		assert.NoError(t, err)

		var count int64
		err = tx.Model(&model.Price{}).Where("item_id = ?", testItemId).Count(&count).Error
		assert.NoError(t, err)
		assert.Equal(t, int64(2), count)

		var deletedPrice model.Price
		err = tx.Where("price_id = ?", price2.PriceId).First(&deletedPrice).Error
		assert.Error(t, err)
		assert.Equal(t, gorm.ErrRecordNotFound, err)
	})
}