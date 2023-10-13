package controllers

import (
	"net/http"
	"password-manager/models"
	"password-manager/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PassManagerInput struct {
	ApplicationName string `json:"application_name" bindings:"required"`
	Password        string `json:"password" bindings:"required"`
}

func CreateNewPassManager(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	userId, err := utils.ExtractTokenId(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input := PassManagerInput{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newData := models.PasswordManager{
		UserID:          userId,
		ApplicationName: input.ApplicationName,
		Password:        input.Password,
	}
	err = newData.Save(db)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Successful Created New Data", "data": newData})
}

func DeletePassManager(ctx *gin.Context) {
	id := ctx.Param("id")
	db := ctx.MustGet("db").(*gorm.DB)
	data := models.PasswordManager{}
	if err := db.Where("id = ?", id).Take(&data).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Delete(&data).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Successful Delete Password Manager Data"})
}

func UpdatePassManager(ctx *gin.Context) {
	id := ctx.Param("id")
	db := ctx.MustGet("db").(*gorm.DB)
	updatePassManagerData := PassManagerInput{}
	if err := ctx.ShouldBindJSON(&updatePassManagerData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	currentPassManagerData := models.PasswordManager{}
	if err := db.Where("id = ?", id).First(&currentPassManagerData).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	encryptedPassword, err := utils.Encrypt(updatePassManagerData.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Model(&currentPassManagerData).Updates(models.PasswordManager{ApplicationName: updatePassManagerData.ApplicationName, Password: encryptedPassword}).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Successful Update Password Manager Data", "data": currentPassManagerData})
}

func GetPasswordManagers(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	data := []models.PasswordManager{}
	if err := db.Find(&data).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, item := range data {
		decryptedPassword, err := utils.Decrypt(item.Password)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		item.Password = decryptedPassword
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Successful Get All Password Managers", "data": data})
}

func GetPasswordManagerDetail(ctx *gin.Context) {
	id := ctx.Param("id")
	db := ctx.MustGet("db").(*gorm.DB)
	data := models.PasswordManager{}
	if err := db.Where("id = ?", id).First(&data).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	decryptedPassword, err := utils.Decrypt(data.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data.Password = decryptedPassword
	ctx.JSON(http.StatusOK, gin.H{"message": "Successful Get Password Manager Detail", "data": data})
}
