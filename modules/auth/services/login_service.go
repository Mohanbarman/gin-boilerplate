package services

import (
	"fmt"
	"net/http"

	"example.com/config"
	dto "example.com/modules/auth/dto"
	models "example.com/modules/users/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func LoginService(c *gin.Context, config *config.Config, db *gorm.DB) {
	loginDto := c.MustGet("data").(*dto.LoginDto)

	user := models.User{}

	fmt.Println(loginDto, user)

	db.Find(&user, &models.User{Email: loginDto.Email})

	if user.ID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"email": []string{"email not found"}})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDto.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Wrong password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"email": user.Email})
}
