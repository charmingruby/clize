package domain

import "time"

type Content struct {
	ID         string    `json:"id"`
	TopicID    string    `json:"topic_id"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}
