package auth

import (
	"example.com/config"
	"example.com/lib"
	"example.com/lib/jwt"
	"example.com/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	Config     *config.Config
	DB         *gorm.DB
	Service    *AuthService
	JwtService *jwt.JwtService
	Smtp       *lib.MailClient
	Redis      *lib.RedisClient
}

func (ctrl *AuthController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		dto := c.MustGet("data").(*LoginDto)

		data, err := ctrl.Service.Login(dto, ctrl.JwtService)

		if err != nil {
			HttpErrors[err.Code].Send(c)
		}

		lib.HttpResponse(200).Data(data).Send(c)
	}
}

func (ctrl *AuthController) GetMe() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("user").(*models.UserModel)
		lib.HttpResponse(200).Data(user.Transform()).Send(c)
	}
}

func (ctrl *AuthController) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		dto := c.MustGet("data").(*RegisterDto)
		result, err := ctrl.Service.Register(dto)

		if err != nil {
			HttpErrors[err.Code].Send(c)
		}

		lib.HttpResponse(200).Data(result).Send(c)
	}
}
