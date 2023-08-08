package router

import (
	"github.com/gin-gonic/gin"
	"sample/handler"
	"sample/middlerware"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	unauthorized := r.Group("v1")
	unauthorized.Use(middlerware.JWTAuth())
	{
		handler.InitRouter(unauthorized)
	}
	return r
}
