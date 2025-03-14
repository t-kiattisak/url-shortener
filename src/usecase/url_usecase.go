package usecase

import (
	"time"
	"url-shortener/src/domain"
	"url-shortener/src/repository"

	"github.com/google/uuid"
)

type ShortenURLUseCase struct {
	Repo repository.URLRepository
}

func NewShortenURLUseCase(repo repository.URLRepository) *ShortenURLUseCase {
	return &ShortenURLUseCase{
		Repo: repo,
	}
}

func (u *ShortenURLUseCase) Execute(originalURL string) (string, error) {
	id := uuid.New().String()
	shortCode := id[:6]
	newURL := domain.URL{
		ID:        id,
		Original:  originalURL,
		Shortened: shortCode,
		Expiry:    time.Now().Add(24 * time.Hour).Unix(),
	}
	err := u.Repo.Save(newURL)
	if err != nil {
		return "", err
	}
	return shortCode, nil
}
