package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/posiposi/project/backend/infrastructure/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("DB connection error: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("DB connection error: %v", err)
	}
	defer sqlDB.Close()

	dbManager := database.NewDBManager(db)
	if err := dbManager.Seed(); err != nil {
		log.Fatalf("Seeder error: %v", err)
	}
}
