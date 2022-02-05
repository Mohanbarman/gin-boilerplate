package auth

import (
	"example.com/config"
	"example.com/lib"
	"example.com/lib/jwt"
	"example.com/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutes(prefix string, rg *gin.RouterGroup, config *config.Config, db *gorm.DB) {
	router := rg.Group(prefix)

	jwtService := jwt.JwtService{Config: &config.Jwt}

	authCtrl := AuthController{
		Config: config,
		DB:     db,
		Service: &AuthService{
			Config: config,
			Db:     db,
		},
		Smtp:       &lib.MailClient{Config: &config.Smtp},
		Redis:      lib.GetRedisClient("reset_password", &config.Redis),
		JwtService: &jwtService,
	}

	authMiddleware := middlewares.AuthMiddleware{
		Jwt: &jwtService,
		DB:  db,
	}

	router.POST("/login", middlewares.Validate(&LoginDto{}), authCtrl.Login())
	router.POST("/register", middlewares.Validate(&RegisterDto{}), authCtrl.Register())
	router.GET("/me", authMiddleware.Validate(jwt.AccessToken), authCtrl.GetMe())
}
