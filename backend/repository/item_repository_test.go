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
	mysqlPort := os.Getenv("MYSQL_TEST_PORT")
	mysqlDatabase := os.Getenv("MYSQL_TEST_DATABASE")
	dsn := fmt.Sprintf("root:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlPassword,
		mysqlHost,
		mysqlPort,
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
	err = seedTestData(db)
	if err != nil {
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
		{UserId: userId1, Name: "User1", Email: "", Password: "password"},
		{UserId: userId2, Name: "User2", Email: "", Password: "password"},
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
		items, err := repo.GetAllItems()
		assert.NoError(t, err)
		assert.Len(t, items, 2)
		assert.IsType(t, domain.Items{}, items)
		assert.IsType(t, domain.Item{}, items[0])
		defer func() {
			db.Exec("DELETE FROM items")
			db.Exec("DELETE FROM users")
		}()
	})
}

func TestGetAllItems_Empty(t *testing.T) {
	t.Run("Get All Items - Empty", func(t *testing.T) {
		db.Exec("DELETE FROM items")
		items, err := repo.GetAllItems()
		assert.NoError(t, err)
		assert.Len(t, items, 0)
	})
}
