package models

import (
	"password-manager/utils"

	"gorm.io/gorm"
)

type PasswordManager struct {
	ID              uint   `json:"id" gorm:"primarykey"`
	UserID          uint   `json:"user_id" gorm:"not null"`
	ApplicationName string `json:"application_name" gorm:"size:255;not null"`
	Password        string `json:"password" gorm:"size:255;not null"`

	// Relationship
	User User `json:"-"`
}

func (p *PasswordManager) Save(db *gorm.DB) error {
	encryptedPassword, err := utils.Encrypt(p.Password)
	if err != nil {
		return err
	}
	p.Password = encryptedPassword
	if err := db.Save(&p).Error; err != nil {
		return err
	}
	return nil
}
