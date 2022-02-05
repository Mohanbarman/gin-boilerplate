package middlewares

import (
	"net/http"
	"strings"

	"example.com/lib"
	"example.com/lib/jwt"
	"example.com/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var unauthorizedErr = gin.H{
	"success": false,
	"message": "Unauthorized",
	"code":    401,
}

type AuthMiddleware struct {
	Jwt *jwt.JwtService
	DB  *gorm.DB
}

func (a *AuthMiddleware) Validate(tokenType jwt.TokenType) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")

		if len(authHeader) == 0 {
			c.JSON(http.StatusUnauthorized, unauthorizedErr)
		}

		token := strings.Split(authHeader, " ")[1]

		if len(token) == 0 {
			c.JSON(http.StatusUnauthorized, unauthorizedErr)
		}

		sub, err := a.Jwt.ParseToken(token, tokenType)

		if err != nil {
			lib.HttpResponse(401).Message("Token is expired").Send(c)
			c.Abort()
			return
		}

		u := models.UserModel{}
		result := a.DB.Find(&u, &models.UserModel{UUID: sub})

		if result.RowsAffected <= 0 {
			lib.HttpResponse(401).Message("User doesn't exists please signup again").Send(c)
			c.Abort()
			return
		}

		c.Set("user", &u)
		c.Next()
	}
}
