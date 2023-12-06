package main

import (
	"log"
	"tgbot/internal/handler"
	"tgbot/internal/repository"
	"tgbot/internal/usecase"

	dotenv "github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := dotenv.Load("../.env"); err != nil {
		log.Println("Cannot find .env file")
	}

	e := echo.New()

	repo := repository.New()
	usecase := usecase.New(repo)
	handler := handler.New(usecase)
	handler.Routes(e)

	if err := e.Start(":3000"); err != nil {
		log.Println(err)
	}

}
