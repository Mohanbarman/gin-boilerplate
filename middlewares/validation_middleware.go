package middlewares

import (
	"example.com/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Validate(s interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.ShouldBind(s); err != nil {
			errors := err.(validator.ValidationErrors)
			c.JSON(400, gin.H{
				"success": false,
				"code":    400,
				"data":    validation.FormatErrors(errors),
			})
			c.Abort()
			return
		}
		c.Set("data", s)
		c.Next()
	}
}
