package middlerware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Token")
		fmt.Println("token: " + token)
		if token == "" {
			//response.Unauthorized(c)
			return
		} else {
			c.Next()
		}
	}
}
