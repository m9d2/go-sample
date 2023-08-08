package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ok(c *gin.Context, data interface{}) {
	res := &Result{
		Code:    http.StatusOK,
		Success: true,
		Message: "ok",
		Data:    data,
	}
	c.JSON(http.StatusOK, res)
}

func Fail(c *gin.Context, err error) {
	res := &Result{
		Code:    http.StatusBadRequest,
		Success: false,
		Message: err.Error(),
	}
	c.JSON(http.StatusOK, res)
}

func Unauthorized(c *gin.Context) {
	res := &Result{
		Code:    http.StatusUnauthorized,
		Success: false,
		Message: "Unauthorized",
	}
	c.JSON(http.StatusOK, res)
}
