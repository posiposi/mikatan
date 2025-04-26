package main

import (
	"fmt"

	"github.com/posiposi/project/backend/db"
	"github.com/posiposi/project/backend/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Book{})
}
