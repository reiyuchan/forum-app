package dto

import "time"

type Comment struct {
	ID        uint      `json:"id"`
	Body      string    `json:"body"`
	User      string    `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdatedComment struct {
	Body      string    `json:"body"`
	UpdatedAt time.Time `json:"updated_at"`
}
