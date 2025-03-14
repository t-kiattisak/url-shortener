package usecase

import (
	"errors"
	"log"
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

func (u *ShortenURLUseCase) ShortenURL(originalURL string) (string, error) {
	id := uuid.New().String()
	shortened := uuid.New().String()[:6]
	expiry := time.Now().Add(7 * 24 * time.Hour)

	newURL := domain.URL{
		ID:        id,
		Original:  originalURL,
		Shortened: shortened,
		Expiry:    expiry,
	}
	err := u.Repo.Save(newURL)
	if err != nil {
		log.Println("Error saving URL:", err)
		return "", err
	}
	return shortened, nil
}

func (u *ShortenURLUseCase) FindByShortened(short string) (*domain.URL, error) {
	url, err := u.Repo.FindByShortened(short)
	if err != nil {
		return nil, errors.New("URL not found")
	}
	return url, nil
}
