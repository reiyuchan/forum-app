package service

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/reiyuchan/forum-app/dto"
	"github.com/reiyuchan/forum-app/util"
	"github.com/reiyuchan/forum-app/util/jwt"
)

var comments []dto.Comment

func CreateComment(ctx *gin.Context) {
	var comment dto.Comment
	err := ctx.ShouldBindJSON(&comment)
	if err != nil || comment.ID == 0 || comment.Body == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	claims := jwt.ExtractClaims(ctx)
	comment.User, err = claims.GetSubject()
	comment.CreatedAt = time.Now()
	if err != nil || comment.User != claims["sub"] {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	comments = append(comments, comment)
	ctx.JSON(http.StatusCreated, comment)
}

func DeleteComment(ctx *gin.Context) {
	id := util.StringToUint(ctx.Query("id"))
	var newComments []dto.Comment
	var deletedComment dto.Comment
	for _, comment := range comments {
		if comment.ID != id {
			newComments = append(newComments, comment)
		}
		if comment.ID == id {
			deletedComment = comment
		}
	}
	comments = newComments
	ctx.JSON(http.StatusOK, gin.H{"message": "Comment deleted...", "comment": deletedComment})
}

func UpdateComment(ctx *gin.Context) {
	id := util.StringToUint(ctx.Query("id"))
	var updatedComment dto.UpdatedComment
	err := ctx.ShouldBindJSON(&updatedComment)
	updatedComment.UpdatedAt = time.Now()
	if err != nil || updatedComment.Body == "" || id == 0 || len(comments) == 0 {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	for i := range len(comments) {
		if comments[i].ID == id {
			comments[i].Body = updatedComment.Body
			ctx.JSON(http.StatusOK, updatedComment)
			return
		} else {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
	}
}
