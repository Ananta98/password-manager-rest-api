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
	auth.POST("/login", controllers.Login)
	auth.POST("/register", controllers.Register)

	passwords := r.Group("/password-managers")
	passwords.Use(middlewares.JwtAuthMiddleware())
	passwords.GET("/", controllers.GetPasswordManagers)
	passwords.GET("/:id", controllers.GetPasswordManagerDetail)
	passwords.POST("/", controllers.CreateNewPassManager)
	passwords.DELETE("/:id", controllers.DeletePassManager)
	passwords.PATCH("/:id", controllers.UpdatePassManager)
	return r
}
