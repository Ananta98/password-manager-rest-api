package models

import (
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Name     string `gorm:"size:255;not null;" json:"name"`
	Email    string `gorm:"size:255;not null;" json:"email"`
	Password string `gorm:"size:255;not null;" json:"-"`
}

func PasswordHashing(passwordText string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passwordText), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func VerifyPassword(passwordInput, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(passwordInput))
}

func (u *User) SaveUser(db *gorm.DB) error {
	hashedPassword, err := PasswordHashing(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	if err := db.Save(&u).Error; err != nil {
		return err
	}
	return nil
}
