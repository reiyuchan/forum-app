package service

import (
	"mal-forums/dto"
	"mal-forums/util/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(ctx *gin.Context) {
	var user dto.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil || user.Username == "" || user.Password == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	tempUser, err := FindOne(&user)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(tempUser.Password), []byte(user.Password))
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	token := jwt.GenerateToken(user.Username)
	ctx.JSON(http.StatusAccepted, gin.H{"user": user.Username, "token": token})
}
