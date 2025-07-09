package repository

import (
	"fmt"
	"log"
	"os"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
	"github.com/posiposi/project/backend/domain"
	"github.com/posiposi/project/backend/internal/orm/model"
	"github.com/stretchr/testify/assert"
)

var repo IItemRepository
var db *gorm.DB

func TestMain(m *testing.M) {
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
	log.Println("Connecting to test database with DSN:", dsn)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to connect to test database:", err)
		os.Exit(1)
	}
	err = db.AutoMigrate(&model.Item{})
	if err != nil {
		os.Exit(1)
	}
	repo = NewItemRepository(db)
	code := m.Run()
	os.Exit(code)
}

func seedTestData(db *gorm.DB) error {
	userID1 := "f47ac10b-58cc-4372-a567-0e02b2c3d111"
	userID2 := "f47ac10b-58cc-4372-a567-0e02b2c3d112"
	users := []model.User{
		{UserId: userID1, Name: "User1", Email: "", Password: "password"},
		{UserId: userID2, Name: "User2", Email: "", Password: "password"},
	}
	items := []model.Item{
		{ItemId: "f47ac10b-58cc-4372-a567-0e02b2c3d110", UserId: userID1, ItemName: "Item1", Stock: true, Description: "Description1"},
		{ItemId: "f47ac10b-58cc-4372-a567-0e02b2c3d111", UserId: userID2, ItemName: "Item2", Stock: false, Description: "Description2"},
	}

	if err := db.Create(&users).Error; err != nil {
		return err
	}

	if err := db.Create(&items).Error; err != nil {
		return err
	}

	return nil
}

func TestGetAllItems(t *testing.T) {
	t.Run("Get All Items - Success", func(t *testing.T) {
		tx := db.Begin()
		defer tx.Rollback()
		err := seedTestData(tx)
		assert.NoError(t, err)
		repo := NewItemRepository(tx)
		items, err := repo.GetAllItems()
		assert.NoError(t, err)
		assert.Len(t, items, 2)
		assert.IsType(t, domain.Items{}, items)
		assert.IsType(t, domain.Item{}, items[0])
	})
}

func TestGetAllItems_Empty(t *testing.T) {
	t.Run("Get All Items - Empty", func(t *testing.T) {
		tx := db.Begin()
		defer tx.Rollback()
		repo := NewItemRepository(tx)
		items, err := repo.GetAllItems()
		assert.NoError(t, err)
		assert.Len(t, items, 0)
	})
}

func TestGetAllItems_SortedByItemId(t *testing.T) {
	t.Run("Get All Items - Sorted by ItemId ASC", func(t *testing.T) {
		tx := db.Begin()
		defer tx.Rollback()

		// Create test data with different item IDs
		userID := "f47ac10b-58cc-4372-a567-0e02b2c3d115"
		user := model.User{UserId: userID, Name: "TestUser", Email: "test@example.com", Password: "password"}
		if err := tx.Create(&user).Error; err != nil {
			t.Fatal(err)
		}

		// Create items with IDs in random order
		items := []model.Item{
			{ItemId: "c47ac10b-58cc-4372-a567-0e02b2c3d003", UserId: userID, ItemName: "Item3", Stock: true, Description: "Desc3"},
			{ItemId: "a47ac10b-58cc-4372-a567-0e02b2c3d001", UserId: userID, ItemName: "Item1", Stock: true, Description: "Desc1"},
			{ItemId: "b47ac10b-58cc-4372-a567-0e02b2c3d002", UserId: userID, ItemName: "Item2", Stock: true, Description: "Desc2"},
		}

		if err := tx.Create(&items).Error; err != nil {
			t.Fatal(err)
		}

		repo := NewItemRepository(tx)
		result, err := repo.GetAllItems()
		assert.NoError(t, err)
		assert.Len(t, result, 3)

		// Check if items are sorted by item_id in ascending order
		assert.Equal(t, "a47ac10b-58cc-4372-a567-0e02b2c3d001", result[0].ItemID())
		assert.Equal(t, "b47ac10b-58cc-4372-a567-0e02b2c3d002", result[1].ItemID())
		assert.Equal(t, "c47ac10b-58cc-4372-a567-0e02b2c3d003", result[2].ItemID())
	})
}

func TestCreateItem(t *testing.T) {
	t.Run("Create Item - Success", func(t *testing.T) {
		tx := db.Begin()
		defer tx.Rollback()

		userID := "f47ac10b-58cc-4372-a567-0e02b2c3d200"
		user := model.User{UserId: userID, Name: "TestUser", Email: "test@example.com", Password: "password"}
		if err := tx.Create(&user).Error; err != nil {
			t.Fatal(err)
		}

		itemID, err := domain.NewItemID("f47ac10b-58cc-4372-a567-0e02b2c3d201")
		assert.NoError(t, err)
		userIDValue, err := domain.NewUserID(userID)
		assert.NoError(t, err)
		itemName, err := domain.NewItemName("Test Item")
		assert.NoError(t, err)
		stock, err := domain.NewStock(true)
		assert.NoError(t, err)
		description, err := domain.NewDescription("Test Description")
		assert.NoError(t, err)
		item, err := domain.NewItem(itemID, *userIDValue, *itemName, *stock, *description)
		assert.NoError(t, err)

		repo := NewItemRepository(tx)
		createdItem, err := repo.CreateItem(item)
		assert.NoError(t, err)
		assert.NotNil(t, createdItem)
		assert.Equal(t, item.ItemID(), createdItem.ItemID())
		assert.Equal(t, item.ItemName(), createdItem.ItemName())
		assert.Equal(t, item.Stock(), createdItem.Stock())
		assert.Equal(t, item.Description(), createdItem.Description())

		var savedItem model.Item
		err = tx.Where("item_id = ?", item.ItemID()).First(&savedItem).Error
		assert.NoError(t, err)
		assert.Equal(t, item.ItemID(), savedItem.ItemId)
		assert.Equal(t, item.ItemName(), savedItem.ItemName)
		assert.Equal(t, item.Stock(), savedItem.Stock)
		assert.Equal(t, item.Description(), savedItem.Description)
	})

	t.Run("Create Item - Duplicate ItemID Error", func(t *testing.T) {
		tx := db.Begin()
		defer tx.Rollback()

		userID := "f47ac10b-58cc-4372-a567-0e02b2c3d300"
		user := model.User{UserId: userID, Name: "TestUser", Email: "test@example.com", Password: "password"}
		if err := tx.Create(&user).Error; err != nil {
			t.Fatal(err)
		}

		itemID := "f47ac10b-58cc-4372-a567-0e02b2c3d301"
		firstItem := model.Item{
			ItemId:      itemID,
			UserId:      userID,
			ItemName:    "First Item",
			Stock:       true,
			Description: "First Description",
		}
		if err := tx.Create(&firstItem).Error; err != nil {
			t.Fatal(err)
		}

		itemIDValue, err := domain.NewItemID(itemID)
		assert.NoError(t, err)
		userIDValue, err := domain.NewUserID(userID)
		assert.NoError(t, err)
		itemName, err := domain.NewItemName("Duplicate Item")
		assert.NoError(t, err)
		stock, err := domain.NewStock(false)
		assert.NoError(t, err)
		description, err := domain.NewDescription("Duplicate Description")
		assert.NoError(t, err)
		duplicateItem, err := domain.NewItem(itemIDValue, *userIDValue, *itemName, *stock, *description)
		assert.NoError(t, err)

		repo := NewItemRepository(tx)
		createdItem, err := repo.CreateItem(duplicateItem)
		assert.Error(t, err)
		assert.Nil(t, createdItem)
	})
}
