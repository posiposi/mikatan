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
	taskRepository := repository.NewBookRepository(db)
	itemRepository := repository.NewItemRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	bookUsecase := usecase.NewBookUsecase(taskRepository)
	itemUsecase := usecase.NewItemUsecase(itemRepository)
	userController := controller.NewUserController(userUsecase)
	bookController := controller.NewBookController(bookUsecase)
	itemController := controller.NewItemController(itemUsecase)
	e := router.NewRouter(userController, bookController, itemController)
	e.Logger.Fatal(e.Start(":8080"))
}
