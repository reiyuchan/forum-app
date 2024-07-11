package middleware

import (
	"mal-forums/util/jwt"
	"net/http"

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
