package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"
	"url-shortener/src/domain"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type URLRepository interface {
	Save(url domain.URL) error
	FindByShortened(short string) (*domain.URL, error)
}

type URLRepositoryPostgres struct {
	DB    *sql.DB
	Redis *redis.Client
}

func (r *URLRepositoryPostgres) Save(url domain.URL) error {
	_, err := r.DB.Exec("INSERT INTO urls (id, original, shortened, expiry) VALUES ($1, $2, $3, $4)",
		url.ID, url.Original, url.Shortened, url.Expiry)

	urlJSON, _ := json.Marshal(url)
	r.Redis.Set(ctx, "url:"+url.Shortened, urlJSON, 7*24*time.Hour)

	return err
}

func (r *URLRepositoryPostgres) FindByShortened(short string) (*domain.URL, error) {
	var url domain.URL

	cachedURL, err := r.Redis.Get(ctx, "url:"+short).Result()
	if err == nil {
		json.Unmarshal([]byte(cachedURL), &url)
		return &url, nil
	}

	err = r.DB.QueryRow(`
        SELECT id, original, shortened, expiry 
        FROM urls 
        WHERE shortened = $1 AND expiry > NOW()`, short).
		Scan(&url.ID, &url.Original, &url.Shortened, &url.Expiry)
	if err != nil {
		return nil, err
	}

	urlJSON, _ := json.Marshal(url)
	r.Redis.Set(ctx, "url:"+short, urlJSON, 7*24*time.Hour)

	return &url, nil
}
