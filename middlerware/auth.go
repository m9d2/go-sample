package middlerware

import (
	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Token")
		if token == "" {
			//response.Unauthorized(c)
			return
		} else {
			c.Next()
		}
	}
}
