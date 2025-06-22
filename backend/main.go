package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/posiposi/project/backend/controller"
	"github.com/posiposi/project/backend/db"
	"github.com/posiposi/project/backend/repository"
	"github.com/posiposi/project/backend/router"
	"github.com/posiposi/project/backend/usecase"
)

func main() {
	db := db.NewDB()
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}
	log.Println("Successfully connected to database")
	userRepository := repository.NewUserRepository(db)
	itemRepository := repository.NewItemRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	itemUsecase := usecase.NewItemUsecase(itemRepository)
	userController := controller.NewUserController(userUsecase)
	itemController := controller.NewItemController(itemUsecase)
	e := router.NewRouter(userController, itemController)
	e.Logger.Fatal(e.Start(":8080"))
}
