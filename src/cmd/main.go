package main

import (
	"url-shortener/src/handler"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	handler := handler.URLHandler{}
	app.Post("/shorten", handler.ShortenURL)
	app.Listen(":3000")
}
