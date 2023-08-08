package model

type Role struct {
	Model
	Name   string `json:"name"`
	Status uint8  `json:"status"`
}

func (Role) TableName() string {
	return "role"
}
