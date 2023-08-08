package router

import (
	"github.com/gin-gonic/gin"
	"sample/app/controller"
	"sample/middlerware"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	unauthorized := r.Group("v1")
	unauthorized.Use(middlerware.JWTAuth())
	{
		controller.InitRouter(unauthorized)
	}
	authorize := r.Group("v1")
	{
		authorize.GET("login", controller.Login)
	}
	return r
}
