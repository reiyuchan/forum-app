package main

import (
	"github.com/reiyuchan/forum-app/router"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	// db.Init()
	router.Routes(&app.RouterGroup)
	app.Run(":8080")
}
