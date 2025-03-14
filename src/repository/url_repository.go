package repository

import "url-shortener/src/domain"

type URLRepository interface {
	Save(url domain.URL) error
	FindByShortened(short string) (*domain.URL, error)
}
