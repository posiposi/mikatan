// Package testutil provides utilities for testing.
package testutil

import (
	"os"
	"testing"

	"github.com/posiposi/project/backend/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupTestDB(t *testing.T) *gorm.DB {
	t.Helper()

	connectionString := os.Getenv("UNIT_TEST_MYSQL_CONNECTION_STRING")
	if connectionString == "" {
		connectionString = "root:test_root_pass@tcp(test_db:3306)/test_db?parseTime=true&loc=Local"
	}

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		t.Fatalf("Failed to migrate test database: %v", err)
	}

	cleanupTestDB(db)

	t.Cleanup(func() {
		cleanupTestDB(db)
	})

	return db
}

func cleanupTestDB(db *gorm.DB) {
	db.Exec("TRUNCATE TABLE users")
}
