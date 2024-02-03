package domain

import "time"

// types: numbered list, simple-list, checklist

type Topic struct {
	ID          string    `json:"id"`
	DocumentID  string    `json:"document_id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	TopicType   string    `json:"topic_type"`
	Contents    []Content `json:"contents"`
	CreatedAt   time.Time `json:"created_at"`
	ModifiedAt  time.Time `json:"modified_at"`
}
