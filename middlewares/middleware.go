package middlewares

import (
	"net/http"
	"password-manager/utils"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := utils.TokenValid(ctx)
		if err != nil {
			ctx.String(http.StatusUnauthorized, err.Error())
			ctx.Abort()
		}
		ctx.Next()
	}
}
