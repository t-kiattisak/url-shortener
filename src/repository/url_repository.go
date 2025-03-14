package repository

import (
	"database/sql"
	"url-shortener/src/domain"
)

type URLRepository interface {
	Save(url domain.URL) error
	FindByShortened(short string) (*domain.URL, error)
}

type URLRepositoryPostgres struct {
	DB *sql.DB
}

func (r *URLRepositoryPostgres) Save(url domain.URL) error {
	_, err := r.DB.Exec("INSERT INTO urls (id, original, shortened, expiry) VALUES ($1, $2, $3, $4)",
		url.ID, url.Original, url.Shortened, url.Expiry)
	return err
}

func (r *URLRepositoryPostgres) FindByShortened(short string) (*domain.URL, error) {
	var url domain.URL
	err := r.DB.QueryRow(`
        SELECT id, original, shortened, expiry 
        FROM urls 
        WHERE shortened = $1 AND expiry > NOW()`, short).
		Scan(&url.ID, &url.Original, &url.Shortened, &url.Expiry)
	if err != nil {
		return nil, err
	}
	return &url, nil
}
