package controller

import (
	"github.com/reiyuchan/forum-app/service"

	"github.com/gin-gonic/gin"
)

func GetPosts(ctx *gin.Context) {
	service.GetPosts(ctx)
}

func SearchPosts(ctx *gin.Context) {
	service.SearchPosts(ctx)
}

func GetUserPosts(ctx *gin.Context) {
	service.GetUserPosts(ctx)
}

func CreatePost(ctx *gin.Context) {
	service.CreatePost(ctx)
}

func DeletePosts(ctx *gin.Context) {
	service.DeletePosts(ctx)
}

func DeletePost(ctx *gin.Context) {
	service.DeletePost(ctx)
}
