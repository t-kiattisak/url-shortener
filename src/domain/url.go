package domain

type URL struct {
	ID        string `json:"id"`
	Original  string `json:"original"`
	Shortened string `json:"shortened"`
	Expiry    int64  `json:"expiry"`
}
