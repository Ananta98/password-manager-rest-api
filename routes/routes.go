package routes

import (
	"password-manager/controllers"
	"password-manager/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
	})
	auth := r.Group("/auth")
	auth.GET("/login", controllers.Login)

	passwords := r.Group("/passwords")
	passwords.Use(middlewares.JwtAuthMiddleware())
	return r
}
