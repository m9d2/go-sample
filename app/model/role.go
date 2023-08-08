package model

import "sample/core/model"

type Role struct {
	model.Model
	Name   string `json:"name"`
	Status uint8  `json:"status"`
}

func (Role) TableName() string {
	return "role"
}

type RoleVO struct {
	Name string `json:"name"`
}
