package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/posiposi/project/backend/controller"
	"github.com/posiposi/project/backend/db"
	"github.com/posiposi/project/backend/infrastructure/openai"
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
	ak := os.Getenv("OPEN_AI_KEY")
	if ak == "" {
		log.Fatalln("OPEN_AI_KEY not set in .env file")
	}
	m := os.Getenv("OPEN_AI_MODEL")
	if m == "" {
		log.Fatalln("OPEN_AI_MODEL not set in .env file")
	}
	h := &http.Client{}
	log.Println("Successfully connected to database")
	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewBookRepository(db)
	OpenAICommunicator := openai.NewOpenAIClient(ak, m, h)
	userUsecase := usecase.NewUserUsecase(userRepository)
	bookUsecase := usecase.NewBookUsecase(taskRepository, OpenAICommunicator)
	userController := controller.NewUserController(userUsecase)
	bookController := controller.NewBookController(bookUsecase)
	e := router.NewRouter(userController, bookController)
	e.Logger.Fatal(e.Start(":8080"))
}
