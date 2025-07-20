package repository

import (
	"fmt"
	"log"
	"os"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/posiposi/project/backend/domain"
	"github.com/posiposi/project/backend/internal/orm/model"
	"github.com/posiposi/project/backend/testutil"
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
	err = testutil.RunPrismaMigrationForTest()
	if err != nil {
		log.Fatalln("Failed to run Prisma migration:", err)
		os.Exit(1)
	}
	repo = NewItemRepository(db)
	code := m.Run()
	os.Exit(code)
}

func seedTestData(db *gorm.DB) error {
	userId1 := "f47ac10b-58cc-4372-a567-0e02b2c3d111"
	userId2 := "f47ac10b-58cc-4372-a567-0e02b2c3d112"
	users := []model.User{
		{Id: userId1, Name: "User1", Email: "user1@example.com", Password: "password", Role: "USER", IsAdmin: false},
		{Id: userId2, Name: "User2", Email: "user2@example.com", Password: "password", Role: "USER", IsAdmin: false},
	}
	items := []model.Item{
		{ItemId: "f47ac10b-58cc-4372-a567-0e02b2c3d110", UserId: userId1, ItemName: "Item1", Stock: true, Description: "Description1"},
		{ItemId: "f47ac10b-58cc-4372-a567-0e02b2c3d111", UserId: userId2, ItemName: "Item2", Stock: false, Description: "Description2"},
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

		userId := "f47ac10b-58cc-4372-a567-0e02b2c3d115"
		user := model.User{Id: userId, Name: "TestUser", Email: "test1@example.com", Password: "password", Role: "USER", IsAdmin: false}
		if err := tx.Create(&user).Error; err != nil {
			t.Fatal(err)
		}

		items := []model.Item{
			{ItemId: "c47ac10b-58cc-4372-a567-0e02b2c3d003", UserId: userId, ItemName: "Item3", Stock: true, Description: "Desc3"},
			{ItemId: "a47ac10b-58cc-4372-a567-0e02b2c3d001", UserId: userId, ItemName: "Item1", Stock: true, Description: "Desc1"},
			{ItemId: "b47ac10b-58cc-4372-a567-0e02b2c3d002", UserId: userId, ItemName: "Item2", Stock: true, Description: "Desc2"},
		}

		if err := tx.Create(&items).Error; err != nil {
			t.Fatal(err)
		}

		repo := NewItemRepository(tx)
		result, err := repo.GetAllItems()
		assert.NoError(t, err)
		assert.Len(t, result, 3)

		// Check if items are sorted by item_id in ascending order
		assert.Equal(t, "a47ac10b-58cc-4372-a567-0e02b2c3d001", result[0].ItemId())
		assert.Equal(t, "b47ac10b-58cc-4372-a567-0e02b2c3d002", result[1].ItemId())
		assert.Equal(t, "c47ac10b-58cc-4372-a567-0e02b2c3d003", result[2].ItemId())
	})
}

func TestCreateItem(t *testing.T) {
	t.Run("Create Item - Success", func(t *testing.T) {
		tx := db.Begin()
		defer tx.Rollback()

		userId := uuid.New().String()
		user := model.User{Id: userId, Name: "TestUser", Email: "test250@example.com", Password: "password", Role: "USER", IsAdmin: false}
		if err := tx.Create(&user).Error; err != nil {
			t.Fatal(err)
		}

		itemId, err := domain.NewItemId(uuid.New().String())
		assert.NoError(t, err)
		userIdValue, err := domain.NewUserId(userId)
		assert.NoError(t, err)
		itemName, err := domain.NewItemName("Test Item")
		assert.NoError(t, err)
		stock, err := domain.NewStock(true)
		assert.NoError(t, err)
		description, err := domain.NewDescription("Test Description")
		assert.NoError(t, err)
		item, err := domain.NewItem(itemId, *userIdValue, *itemName, *stock, *description)
		assert.NoError(t, err)

		repo := NewItemRepository(tx)
		createdItem, err := repo.CreateItem(item)
		assert.NoError(t, err)
		assert.NotNil(t, createdItem)
		assert.Equal(t, item.ItemId(), createdItem.ItemId())
		assert.Equal(t, item.ItemName(), createdItem.ItemName())
		assert.Equal(t, item.Stock(), createdItem.Stock())
		assert.Equal(t, item.Description(), createdItem.Description())

		var savedItem model.Item
		err = tx.Where("item_id = ?", item.ItemId()).First(&savedItem).Error
		assert.NoError(t, err)
		assert.Equal(t, item.ItemId(), savedItem.ItemId)
		assert.Equal(t, item.ItemName(), savedItem.ItemName)
		assert.Equal(t, item.Stock(), savedItem.Stock)
		assert.Equal(t, item.Description(), savedItem.Description)
	})

	t.Run("Create Item - Duplicate ItemId Error", func(t *testing.T) {
		tx := db.Begin()
		defer tx.Rollback()

		userId := uuid.New().String()
		user := model.User{Id: userId, Name: "TestUser", Email: "test3@example.com", Password: "password", Role: "USER", IsAdmin: false}
		if err := tx.Create(&user).Error; err != nil {
			t.Fatal(err)
		}

		itemId := uuid.New().String()
		firstItem := model.Item{
			ItemId:      itemId,
			UserId:      userId,
			ItemName:    "First Item",
			Stock:       true,
			Description: "First Description",
		}
		if err := tx.Create(&firstItem).Error; err != nil {
			t.Fatal(err)
		}

		itemIdValue, err := domain.NewItemId(itemId)
		assert.NoError(t, err)
		userIdValue, err := domain.NewUserId(userId)
		assert.NoError(t, err)
		itemName, err := domain.NewItemName("Duplicate Item")
		assert.NoError(t, err)
		stock, err := domain.NewStock(false)
		assert.NoError(t, err)
		description, err := domain.NewDescription("Duplicate Description")
		assert.NoError(t, err)
		duplicateItem, err := domain.NewItem(itemIdValue, *userIdValue, *itemName, *stock, *description)
		assert.NoError(t, err)

		repo := NewItemRepository(tx)
		createdItem, err := repo.CreateItem(duplicateItem)
		assert.Error(t, err)
		assert.Nil(t, createdItem)
	})
}

func TestGetItemByID(t *testing.T) {
	t.Run("Get Item By ID - Success", func(t *testing.T) {
		tx := db.Begin()
		defer tx.Rollback()

		userId := "f47ac10b-58cc-4372-a567-0e02b2c3d400"
		user := model.User{Id: userId, Name: "TestUser", Email: "test400@example.com", Password: "password", Role: "USER", IsAdmin: false}
		if err := tx.Create(&user).Error; err != nil {
			t.Fatal(err)
		}

		itemId := "f47ac10b-58cc-4372-a567-0e02b2c3d401"
		item := model.Item{
			ItemId:      itemId,
			UserId:      userId,
			ItemName:    "Test Item",
			Stock:       true,
			Description: "Test Description",
		}
		if err := tx.Create(&item).Error; err != nil {
			t.Fatal(err)
		}

		itemIdValue, err := domain.NewItemId(itemId)
		assert.NoError(t, err)

		repo := NewItemRepository(tx)
		result, err := repo.GetItemByID(itemIdValue)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, itemId, result.ItemId())
		assert.Equal(t, "Test Item", result.ItemName())
		assert.Equal(t, true, result.Stock())
		assert.Equal(t, "Test Description", result.Description())
	})

	t.Run("Get Item By ID - Not Found", func(t *testing.T) {
		tx := db.Begin()
		defer tx.Rollback()

		itemIdValue, err := domain.NewItemId("f47ac10b-58cc-4372-a567-0e02b2c3d999")
		assert.NoError(t, err)

		repo := NewItemRepository(tx)
		result, err := repo.GetItemByID(itemIdValue)
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestUpdateItem(t *testing.T) {
	t.Run("Update Item - Success", func(t *testing.T) {
		tx := db.Begin()
		defer tx.Rollback()

		userId := "f47ac10b-58cc-4372-a567-0e02b2c3d500"
		user := model.User{Id: userId, Name: "TestUser", Email: "test500@example.com", Password: "password", Role: "USER", IsAdmin: false}
		if err := tx.Create(&user).Error; err != nil {
			t.Fatal(err)
		}

		itemId := "f47ac10b-58cc-4372-a567-0e02b2c3d501"
		originalItem := model.Item{
			ItemId:      itemId,
			UserId:      userId,
			ItemName:    "Original Item",
			Stock:       true,
			Description: "Original Description",
		}
		if err := tx.Create(&originalItem).Error; err != nil {
			t.Fatal(err)
		}

		itemIdValue, err := domain.NewItemId(itemId)
		assert.NoError(t, err)
		userIdValue, err := domain.NewUserId(userId)
		assert.NoError(t, err)
		itemName, err := domain.NewItemName("Updated Item")
		assert.NoError(t, err)
		stock, err := domain.NewStock(false)
		assert.NoError(t, err)
		description, err := domain.NewDescription("Updated Description")
		assert.NoError(t, err)
		updatedItem, err := domain.NewItem(itemIdValue, *userIdValue, *itemName, *stock, *description)
		assert.NoError(t, err)

		repo := NewItemRepository(tx)
		result, err := repo.UpdateItem(updatedItem)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, itemId, result.ItemId())
		assert.Equal(t, "Updated Item", result.ItemName())
		assert.Equal(t, false, result.Stock())
		assert.Equal(t, "Updated Description", result.Description())

		var savedItem model.Item
		err = tx.Where("item_id = ?", itemId).First(&savedItem).Error
		assert.NoError(t, err)
		assert.Equal(t, "Updated Item", savedItem.ItemName)
		assert.Equal(t, false, savedItem.Stock)
		assert.Equal(t, "Updated Description", savedItem.Description)
	})

	t.Run("Update Item - Not Found", func(t *testing.T) {
		tx := db.Begin()
		defer tx.Rollback()

		itemIdValue, err := domain.NewItemId(uuid.New().String())
		assert.NoError(t, err)
		userIdValue, err := domain.NewUserId(uuid.New().String())
		assert.NoError(t, err)
		itemName, err := domain.NewItemName("Non-existent Item")
		assert.NoError(t, err)
		stock, err := domain.NewStock(true)
		assert.NoError(t, err)
		description, err := domain.NewDescription("Non-existent Description")
		assert.NoError(t, err)
		item, err := domain.NewItem(itemIdValue, *userIdValue, *itemName, *stock, *description)
		assert.NoError(t, err)

		repo := NewItemRepository(tx)
		result, err := repo.UpdateItem(item)
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestDeleteItem(t *testing.T) {
	t.Run("Delete Item - Success", func(t *testing.T) {
		tx := db.Begin()
		defer tx.Rollback()

		userId := uuid.New().String()
		user := model.User{Id: userId, Name: "TestUser", Email: "test600@example.com", Password: "password", Role: "USER", IsAdmin: false}
		if err := tx.Create(&user).Error; err != nil {
			t.Fatal(err)
		}

		itemId := uuid.New().String()
		item := model.Item{
			ItemId:      itemId,
			UserId:      userId,
			ItemName:    "Item to Delete",
			Stock:       true,
			Description: "Description to Delete",
		}
		if err := tx.Create(&item).Error; err != nil {
			t.Fatal(err)
		}

		itemIdValue, err := domain.NewItemId(itemId)
		assert.NoError(t, err)

		repo := NewItemRepository(tx)
		err = repo.DeleteItem(itemIdValue)
		assert.NoError(t, err)

		var deletedItem model.Item
		err = tx.Where("item_id = ?", itemId).First(&deletedItem).Error
		assert.Error(t, err)
		assert.True(t, gorm.ErrRecordNotFound == err)
	})

	t.Run("Delete Item - Not Found", func(t *testing.T) {
		tx := db.Begin()
		defer tx.Rollback()

		itemIdValue, err := domain.NewItemId(uuid.New().String())
		assert.NoError(t, err)

		repo := NewItemRepository(tx)
		err = repo.DeleteItem(itemIdValue)
		assert.Error(t, err)
	})
}
