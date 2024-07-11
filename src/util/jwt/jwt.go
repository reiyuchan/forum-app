package jwt

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("tokyoKen")

func GenerateToken(username string) string {
	claims := jwt.MapClaims{"sub": username}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return signedToken
}

func VerifyToken(tkn string) *jwt.Token {
	token, err := jwt.Parse(tkn, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	if !token.Valid {
		fmt.Println("Invalid token...")
		return nil
	}
	return token
}

func ExtractToken(ctx *gin.Context) string {
	bearerToken := ctx.Request.Header.Get("Authorization")
	splittedBearerToken := strings.Split(bearerToken, " ")
	if len(splittedBearerToken) == 2 && splittedBearerToken[0] == "Bearer" {
		return splittedBearerToken[1]
	}
	return ""
}

func ExtractClaims(ctx *gin.Context) jwt.MapClaims {
	token := ExtractToken(ctx)
	if token == "" {
		fmt.Println("Token is empty...")
		return nil
	}
	claims := VerifyToken(token).Claims.(jwt.MapClaims)
	return claims
}
