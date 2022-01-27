package services

import (
	"net/http"

	"example.com/config"
	dto "example.com/modules/auth/dto"
	models "example.com/modules/users/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RegisterService(c *gin.Context, config *config.Config, dbClient *gorm.DB) {
	registerDto := c.MustGet("data").(*dto.RegisterDto)

	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(registerDto.Password), 10)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to hash password"})
		return
	}

	if created := dbClient.Create(&models.User{
		Email:    registerDto.Email,
		Password: string(hashedPasswordBytes),
	}); created.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"email": []string{"Email already exists"}})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"email": registerDto.Email,
	})

}
