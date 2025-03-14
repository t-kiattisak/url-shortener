package main

import (
	"log"
	"url-shortener/src/handler"
	"url-shortener/src/infrastructure"
	"url-shortener/src/repository"
	"url-shortener/src/usecase"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	db := infrastructure.NewPostgresDB()
	redis := infrastructure.NewRedisClient()

	repo := &repository.URLRepositoryPostgres{DB: db,
		Redis: redis,
	}

	shortenUseCase := usecase.NewShortenURLUseCase(repo)
	handler := handler.NewURLHandler(shortenUseCase)

	app.Post("/shorten", handler.ShortenURL)
	app.Get("/:shortened", handler.RedirectURL)

	log.Println("Service is running on port 3000...")
	log.Fatal(app.Listen(":3000"))
}
