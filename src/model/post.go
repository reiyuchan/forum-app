package model

import (
	"time"

	"gorm.io/gorm"
)

type post struct {
	gorm.Model
	ID        uint
	Title     string
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
