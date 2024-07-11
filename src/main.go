package main

import (
	"mal-forums/router"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	// db.Init()
	router.Routes(&app.RouterGroup)
	app.Run(":8080")
}
