package models

import "time"

type ChatMessage struct {
	ID        uint      `json:"id"`
	Message   string    `json:"message"`
	From      uint      `json:"message_from" db:"message_from"`
	To        uint      `json:"message_to" db:"message_to"`
	IsEdited  bool      `json:"is_edited" db:"is_edited"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
