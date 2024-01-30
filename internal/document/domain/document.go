package domain

import "time"

type Document struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Pretension  string    `json:"pretension"`
	Topics      []Topic   `json:"topics"`
	CreatedAt   time.Time `json:"created_at"`
	ModifiedAt  time.Time `json:"modified_at"`
}
