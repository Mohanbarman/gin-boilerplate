package db

import (
	user_models "example.com/modules/users/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&user_models.User{})
}
