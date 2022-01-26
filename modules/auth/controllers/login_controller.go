package auth

import (
	"example.com/config"
	services "example.com/modules/auth/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LoginController(config *config.Config, dbClient *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		services.LoginService(c, config, dbClient)
	}
}
