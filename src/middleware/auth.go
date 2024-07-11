package middleware

import (
	"net/http"

	"github.com/reiyuchan/forum-app/util/jwt"

	"github.com/gin-gonic/gin"
)

func Authorize(ctx *gin.Context) {
	token := jwt.ExtractToken(ctx)
	if token == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	tkn := jwt.VerifyToken(token)
	if tkn == nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	ctx.Next()
}
