package controller

import (
	"mal-forums/service"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	service.Login(ctx)
}
