package controller

import (
	"sample/app/model"
	"sample/app/query"
	"sample/app/service"
	"sample/core/response"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserService
}

func InitRouter(g *gin.RouterGroup) {
	u := UserController{}
	g.GET("user", u.findAll)
	g.POST("user", u.save)
	g.PUT("user", u.update)
	g.DELETE("user", u.delete)
}

func (u *UserController) findAll(c *gin.Context) {
	var userQuery query.User
	// if err := c.ShouldBindJSON(&userQuery); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	users, err := u.userService.FindAllUser(&userQuery)
	if err != nil {
		response.Fail(c, err)
	} else {
		response.Ok(c, users)
	}
}

func (u *UserController) save(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.Fail(c, err)
	}
	e := u.userService.SaveUser(&user)
	if e != nil {
		response.Fail(c, e)
	} else {
		response.Ok(c, nil)
	}
}

func (u *UserController) update(c *gin.Context) {

}

func (u *UserController) delete(c *gin.Context) {

}
