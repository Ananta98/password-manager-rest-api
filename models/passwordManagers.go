package models

import "gorm.io/gorm"

type PasswordManagers struct {
	ID              uint   `json:"id" gorm:"primarykey"`
	UserID          uint   `json:"user_id" gorm:"not null"`
	ApplicationName string `json:"application_name" gorm:"size:255;not null"`
	Password        string `json:"-" gorm:"size:255;not null"`

	// Relationship
	User User `json:"-"`
}

func (p *PasswordManagers) Save(db *gorm.DB) error {
	encryptedPassword := p.Password
	p.Password = encryptedPassword
	if err := db.Save(&p).Error; err != nil {
		return err
	}
	return nil
}
