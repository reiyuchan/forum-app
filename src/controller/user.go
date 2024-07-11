package controller

import (
	"mal-forums/service"

	"github.com/gin-gonic/gin"
)

func GetUsers(ctx *gin.Context) {
	service.GetUsers(ctx)
}

func GetUser(ctx *gin.Context) {
	service.GetUser(ctx)
}

func CreateUser(ctx *gin.Context) {
	service.CreateUser(ctx)
}

func DeleteUsers(ctx *gin.Context) {
	service.DeleteUsers(ctx)
}

func DeleteUser(ctx *gin.Context) {
	service.DeleteUser(ctx)
}
