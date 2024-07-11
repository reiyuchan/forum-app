package dto

import "time"

type Post struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	User      string    `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
