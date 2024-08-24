package router

import (
	"github.com/reiyuchan/forum-app/controller"
	"github.com/reiyuchan/forum-app/middleware"

	"github.com/gin-gonic/gin"
)

func Routes(routes *gin.RouterGroup) {
	const POST = "/post"
	const USER = "/user"
	const COMMENT = "/comment"

	api := routes.Group("/api/v1")
	{
		posts := api.Group(POST)
		{
			auth := posts.Group("")
			auth.Use(middleware.Authorize)
			{
				auth.POST("/create", controller.CreatePost)
				auth.DELETE("", controller.DeletePost)
				auth.GET("", controller.GetUserPosts)
				auth.PUT("/update", controller.UpdatePost)
			}
			posts.GET("/search", controller.SearchPosts)
			posts.GET("/index", controller.GetPosts)
			posts.DELETE("/index", controller.DeletePosts)
		}

		comments := api.Group(POST + COMMENT)
		comments.Use(middleware.Authorize)
		{
			comments.POST("/create", controller.CreateComment)
			comments.DELETE("", controller.DeleteComment)
			comments.PUT("/update", controller.UpdateComment)
		}

		users := api.Group(USER)
		{
			auth := users.Group("")
			auth.Use(middleware.Authorize)
			{
				auth.GET("", controller.GetUser)
				auth.DELETE("", controller.DeleteUser)
			}
			users.POST("/login", controller.Login)
			users.POST("/create", controller.CreateUser)
			users.GET("/index", controller.GetUsers)
			users.DELETE("/index", controller.DeleteUsers)
		}
	}
}
