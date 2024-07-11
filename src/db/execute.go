package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Current *gorm.DB

const (
	HOST     = "localhost"
	PORT     = 5432
	USER     = "postgres"
	PASSWORD = "postgres"
	DBNAME   = "forums"
	SSLMODE  = "disable"
)

func execute() {
	connectionUri := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", HOST, USER, PASSWORD, DBNAME, PORT, SSLMODE)

	db, err := gorm.Open(postgres.Open(connectionUri), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	Current = db
}
