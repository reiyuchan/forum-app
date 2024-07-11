package service

import (
	"fmt"
	"mal-forums/dto"
	"mal-forums/util"
	"mal-forums/util/jwt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var users = []dto.User{}

func GetUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, users)
}

func GetUser(ctx *gin.Context) {
	claims := jwt.ExtractClaims(ctx)
	for _, user := range users {
		if user.Username == claims["sub"] || user.Email == claims["sub"] {
			userData := dto.UserData{ID: user.ID, Username: user.Username, Email: user.Email, CreatedAt: user.CreatedAt}
			ctx.JSON(http.StatusOK, userData)
			return
		}
	}
	ctx.AbortWithStatus(http.StatusNotFound)
}

func CreateUser(ctx *gin.Context) {
	var user dto.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil || user.Username == "" || user.Email == "" || user.Password == "" || user.Password_Confirmation == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.CreatedAt = time.Now()
	isUserExist := userExists(&user)
	if err != nil || isUserExist {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	user.Password = string(hashedPass)
	user.Password_Confirmation = string(hashedPass)
	userData := dto.UserData{ID: user.ID, Username: user.Username, Email: user.Email, CreatedAt: user.CreatedAt}
	users = append(users, user)
	ctx.JSON(http.StatusCreated, userData)
}

func DeleteUsers(ctx *gin.Context) {
	var newUsers []dto.User
	users = newUsers
	ctx.JSON(http.StatusOK, users)
}

func DeleteUser(ctx *gin.Context) {
	id := util.StringToUint(ctx.Query("id"))
	var newUsers []dto.User
	var deletedUser dto.User
	for _, user := range users {
		if user.ID != id {
			newUsers = append(newUsers, user)
		}
		if user.ID == id {
			deletedUser = user
		}
	}
	users = newUsers
	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted..", "user": deletedUser})
}

func FindOne(user *dto.User) (*dto.User, error) {
	for _, storedUser := range users {
		if user.Username == storedUser.Email || user.Username == storedUser.Username || user.Email == storedUser.Email {
			return &storedUser, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}

func userExists(user *dto.User) bool {
	storedUser, err := FindOne(user)
	if err != nil {
		return false
	}
	if storedUser != nil {
		return true
	}
	return false
}
