package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var API_SECRET = GetEnv("API_SECRET", "rahasiasekali")

func GenerateToken(userId uint) (string, error) {
	tokenLifeSpan, err := strconv.Atoi(GetEnv("TOKEN_HOUR_LIFESPAN", "1"))
	if err != nil {
		return "", err
	}
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userId
	claims["exp"] = time.Now().Add(time.Duration(tokenLifeSpan) * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(API_SECRET))
}

func GetToken(ctx *gin.Context) string {
	token := ctx.Query("token")
	if token != "" {
		return token
	}
	header := ctx.Request.Header.Get("Authorization")
	splitted := strings.Split(header, " ")
	if len(splitted) == 2 {
		return splitted[1]
	}
	return ""
}

func TokenValid(ctx *gin.Context) error {
	token := GetToken(ctx)
	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(API_SECRET), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func ExtractTokenId(ctx *gin.Context) (uint, error) {
	token := GetToken(ctx)
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(API_SECRET), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if jwtToken.Valid && ok {
		uid, err := strconv.ParseUint(fmt.Sprintf("%0.f", claims["user_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint(uid), nil
	}
	return 0, nil
}
