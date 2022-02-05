package models

import (
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	UUID           string `gorm:"unique"`
	Email          string `gorm:"email;unique;not null"`
	Password       string
	Name           string
	Status         string
	ProfilePicture string
}

func (user *UserModel) BeforeCreate(scope *gorm.DB) (err error) {
	user.UUID = uuid.NewV4().String()
	return
}

func (user *UserModel) TableName() string {
	return "users"
}

func (user *UserModel) Transform() map[string]interface{} {
	return map[string]interface{}{
		"id":         user.UUID,
		"email":      user.Email,
		"name":       user.Name,
		"profile":    user.ProfilePicture,
		"status":     user.Status,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
	}
}

func (user *UserModel) SetPassword(password string) (err error) {
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	user.Password = string(hashedPasswordBytes)
	return
}
