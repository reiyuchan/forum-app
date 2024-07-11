package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID                    uint
	Email                 string
	Username              string
	Password              string
	Password_Confirmation string
	CreatedAt             time.Time
}
