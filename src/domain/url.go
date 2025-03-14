package domain

import "time"

type URL struct {
	ID        string    `json:"id"`
	Original  string    `json:"original"`
	Shortened string    `json:"shortened"`
	Expiry    time.Time `json:"expiry"`
}
