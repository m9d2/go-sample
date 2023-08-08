package query

import "sample/core/model"

type User struct {
	model.Query
	Name string `json:"name" gorm:"size:255"`
}
