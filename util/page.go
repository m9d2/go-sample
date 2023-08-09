package util

type PageReq struct {
	Page int `form:"page"`
	Size int `form:"size"`
}
