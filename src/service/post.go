package service

import (
	"net/http"
	"time"

	"github.com/reiyuchan/forum-app/dto"
	"github.com/reiyuchan/forum-app/util"
	"github.com/reiyuchan/forum-app/util/jwt"

	"github.com/gin-gonic/gin"
)

var posts = []dto.Post{}

func GetPosts(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, posts)
}

func SearchPosts(ctx *gin.Context) {
	query := ctx.Query("query")
	user := ctx.Query("user")
	if query == "" && user == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	var searchedPosts []dto.Post
	for _, post := range posts {
		if util.Match(query, post.Title) && util.Match(user, post.User) {
			searchedPosts = append(searchedPosts, post)
		}
	}
	if len(searchedPosts) == 0 {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	ctx.JSON(http.StatusOK, searchedPosts)
}

func GetUserPosts(ctx *gin.Context) {
	claims := jwt.ExtractClaims(ctx)
	var userPosts []dto.Post
	for _, post := range posts {
		if post.User == claims["sub"] {
			userPosts = append(userPosts, post)
		}
	}
	if len(userPosts) == 0 {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	ctx.JSON(http.StatusOK, userPosts)
}

func CreatePost(ctx *gin.Context) {
	var post dto.Post
	err := ctx.ShouldBindJSON(&post)
	if err != nil || post.ID == 0 || post.Title == "" || post.Body == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	claims := jwt.ExtractClaims(ctx)
	post.User, err = claims.GetSubject()
	post.CreatedAt = time.Now()
	if err != nil || post.User != claims["sub"] {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	posts = append(posts, post)
	ctx.JSON(http.StatusOK, post)
}

func DeletePosts(ctx *gin.Context) {
	var newPosts []dto.Post
	posts = newPosts
	ctx.JSON(http.StatusOK, posts)
}

func DeletePost(ctx *gin.Context) {
	id := util.StringToUint(ctx.Query("id"))
	var newPosts []dto.Post
	var deletedPost dto.Post
	for _, post := range posts {
		if post.ID != id {
			newPosts = append(newPosts, post)
		}
		if post.ID == id {
			deletedPost = post
		}
	}
	posts = newPosts
	ctx.JSON(http.StatusOK, gin.H{"message": "Post deleted...", "post": deletedPost})
}

func UpdatePost(ctx *gin.Context) {
	id := util.StringToUint(ctx.Query("id"))
	var updatedPost dto.UpdatedPost
	err := ctx.ShouldBindJSON(updatedPost)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	for i := range len(posts) {
		if posts[i].ID == id {
			posts[i].Body = updatedPost.Body
		}

	}
	ctx.JSON(http.StatusOK, updatedPost)
}
