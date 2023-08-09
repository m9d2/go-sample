package vo

import (
	"sample/model"
	"sample/util"
)

type FindAllUserReq struct {
	util.PageReq
	Name string `form:"name"`
}

type FindAllUserRsp struct {
	model.Model
	Name        string `json:"name"`
	DeptId      uint   `json:"dept_id"`
	Username    string `json:"username"`
	Nickname    string `json:"nickname"`
	UserType    uint8  `json:"user_type"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Sex         uint8  `json:"sex"`
	Avatar      string `json:"avatar"`
	Status      uint8  `json:"status"`
	LastLoginIp string `json:"last_login_ip"`
	Remark      string `json:"remark"`
	RoleName    string `json:"roleName"`
}

type SaveUserReq struct {
	Name     string `json:"name"`
	DeptId   uint   `json:"dept_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	UserType uint8  `json:"user_type"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Sex      uint8  `json:"sex"`
	Avatar   string `json:"avatar"`
	Status   uint8  `json:"status"`
	Remark   string `json:"remark"`
	RoleId   uint   `json:"roleId"`
}
