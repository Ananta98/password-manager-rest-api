package config

import (
	"fmt"
	"password-manager/models"
	"password-manager/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	username := utils.GetEnv("DATABASE_USERNAME", "root")
	password := utils.GetEnv("DATABASE_PASSWORD", "root")
	host := utils.GetEnv("DATABASE_HOST", "127.0.0.1")
	port := utils.GetEnv("DATABASE_PORT", "8080")
	database_name := utils.GetEnv("DATABASE_NAME", "db_password_manager")
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database_name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(models.User{})
	return db
}
