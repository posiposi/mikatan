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
	"github.com/stretchr/testify/assert"
)

var repo IItemRepository
var db *gorm.DB

func TestMain(m *testing.M) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalln("Error loading .env file")
	}
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlPort := os.Getenv("MYSQL_PORT")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")
	dsn := fmt.Sprintf("%s:%s@tcp(db:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlUser,
		mysqlPassword,
		mysqlPort,
		mysqlDatabase,
	)
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
	userId1 := uuid.NewString()
	userId2 := uuid.NewString()
	users := []model.User{
		{UserId: userId1, Name: "User1", Email: "", Password: "password"},
		{UserId: userId2, Name: "User2", Email: "", Password: "password"},
	}
	items := []model.Item{
		{ItemId: uuid.NewString(), UserId: userId1, ItemName: "Item1", Stock: true, Description: "Description1"},
		{ItemId: uuid.NewString(), UserId: userId2, ItemName: "Item2", Stock: false, Description: "Description2"},
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
