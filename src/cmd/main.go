package main

import (
	"url-shortener/src/handler"
	"url-shortener/src/infrastructure"
	"url-shortener/src/repository"
	"url-shortener/src/usecase"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	db := infrastructure.NewPostgresDB()
	repo := &repository.URLRepositoryPostgres{DB: db}
	shortenUseCase := usecase.NewShortenURLUseCase(repo)
	handler := handler.NewURLHandler(shortenUseCase)
	app.Post("/shorten", handler.ShortenURL)
	app.Listen(":3000")
}
