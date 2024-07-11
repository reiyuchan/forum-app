package controller

import (
	"github.com/reiyuchan/forum-app/service"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	service.Login(ctx)
}
