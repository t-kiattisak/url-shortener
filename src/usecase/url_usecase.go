package usecase

import (
	"time"
	"url-shortener/src/domain"
	"url-shortener/src/repository"

	"github.com/google/uuid"
)

type ShortenedURLUsecase struct {
	repo repository.URLRepository
}

func (u *ShortenedURLUsecase) Execute(originalURL string) (string, error) {
	id := uuid.New().String()
	shortCode := id[:6]
	newURL := domain.URL{
		ID:        id,
		Original:  originalURL,
		Shortened: shortCode,
		Expiry:    time.Now().Add(24 * time.Hour).Unix(),
	}
	err := u.repo.Save(newURL)
	if err != nil {
		return "", err
	}
	return shortCode, nil
}
