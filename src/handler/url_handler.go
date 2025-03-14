package handler

import (
	"url-shortener/src/usecase"

	"github.com/gofiber/fiber/v2"
)

type URLHandler struct {
	ShortenUseCase *usecase.ShortenURLUseCase
}

func NewURLHandler(shortenUseCase *usecase.ShortenURLUseCase) *URLHandler {
	return &URLHandler{ShortenUseCase: shortenUseCase}
}

func (h *URLHandler) ShortenURL(c *fiber.Ctx) error {
	var req struct {
		URL string `json:"url"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	short, err := h.ShortenUseCase.ShortenURL(req.URL)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to shorten URL"})
	}

	return c.JSON(fiber.Map{"short_url": "http://localhost:3000/" + short})
}

func (h *URLHandler) RedirectURL(c *fiber.Ctx) error {
	shortened := c.Params("shortened")
	url, err := h.ShortenUseCase.FindByShortened(shortened)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "URL not found",
		})
	}
	return c.Redirect(url.Original, fiber.StatusFound)
}
