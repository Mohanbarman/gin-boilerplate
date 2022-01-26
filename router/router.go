package router

import (
	"example.com/config"
	"example.com/modules/auth"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// all route groups should sit here
func SetupRoutes(engine *gin.Engine, config *config.Config, db *gorm.DB) {
	rg := engine.Group("/api")

	auth.InitRoutes("/auth", rg, config, db)
}
