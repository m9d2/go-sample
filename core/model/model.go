package model

type Model struct {
	ID        uint `json:"id" gorm:"primarykey"`
	CreatedAt Time `json:"createdAt"`
	UpdateAt  Time `json:"updateAt"`
}
