package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"sample/core/response"
)

func Login(c *gin.Context) {
	err := errors.New("这个是错误")
	response.Fail(c, err)
}
