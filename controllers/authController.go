package controllers

import (
	"net/http"
	"password-manager/models"
	"password-manager/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type InputLoginAuth struct {
	InputUser string `json:"input_user" bindings:"required"`
	Password  string `json:"password" bindings:"required"`
}

type RegisterUserInput struct {
	Username             string `bindings:"required" json:"username"`
	Name                 string `bindings:"required" json:"name"`
	Email                string `bindings:"required" json:"email"`
	Password             string `bindings:"required" json:"password"`
	PasswordConfirmation string `bindings:"required" json:"password_confirmation"`
}

func Login(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	input := InputLoginAuth{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	foundUser := models.User{}
	if err := db.Where("username = ? or email = ?", input.InputUser, input.InputUser).Take(&foundUser).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	err := models.VerifyPassword(input.Password, foundUser.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	token, err := utils.GenerateToken(foundUser.ID)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Successful Login", "token": token})
}

func Register(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	input := RegisterUserInput{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if input.Password != input.PasswordConfirmation {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Make Sure Pasword Input and Password Confirmation should same"})
		return
	}
	newUser := models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
		Name:     input.Name,
	}
	if err := newUser.SaveUser(db); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Successful Register", "user": newUser})
}
