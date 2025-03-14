package handler

import (
	"url-shortener/src/usecase"

	"github.com/gofiber/fiber/v2"
)

type URLHandler struct {
	shortenUseCase usecase.ShortenedURLUsecase
}

func (h *URLHandler) ShortenURL(c *fiber.Ctx) error {
	var req struct {
		URL string `json:"url"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	short, err := h.shortenUseCase.Execute(req.URL)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to shorten URL"})
	}

	return c.JSON(fiber.Map{"short_url": "http://localhost:3000/" + short})
}
