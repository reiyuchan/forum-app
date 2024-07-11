package dto

import "time"

type User struct {
	ID                    uint      `json:"id"`
	Email                 string    `json:"email"`
	Username              string    `json:"username"`
	Password              string    `json:"password"`
	Password_Confirmation string    `json:"password_confirmation"`
	CreatedAt             time.Time `json:"created_at"`
}

type UserData struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}
