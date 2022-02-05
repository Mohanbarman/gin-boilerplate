package db

import (
	"example.com/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	m := []interface{}{
		&models.UserModel{},
	}
	db.AutoMigrate(m...)
}
