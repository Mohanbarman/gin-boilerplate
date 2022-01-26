package auth

import (
	"example.com/config"
	"example.com/middlewares"
	controllers "example.com/modules/auth/controllers"
	dto "example.com/modules/auth/dto"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// all routes of the module
func InitRoutes(prefix string, rg *gin.RouterGroup, config *config.Config, db *gorm.DB) {
	router := rg.Group(prefix)

	router.POST("/login", middlewares.Validate(&dto.LoginDto{}), controllers.LoginController(config, db))
	router.POST("/register", middlewares.Validate(&dto.RegisterDto{}), controllers.RegisterController(config, db))
}
