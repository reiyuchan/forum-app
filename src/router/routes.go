package router

import (
	"mal-forums/controller"
	"mal-forums/middleware"

	"github.com/gin-gonic/gin"
)

func Routes(routes *gin.RouterGroup) {
	const POST = "/post"
	const USER = "/user"

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
			}
			posts.GET("/search", controller.SearchPosts)
			posts.PUT("/:id")
			posts.GET("/index", controller.GetPosts)
			posts.DELETE("/index", controller.DeletePosts)
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
