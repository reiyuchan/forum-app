package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/reiyuchan/forum-app/service"
)

func CreateComment(ctx *gin.Context) {
	service.CreateComment(ctx)
}

func DeleteComment(ctx *gin.Context) {
	service.DeleteComment(ctx)
}

func UpdateComment(ctx *gin.Context) {
	service.UpdateComment(ctx)
}
