package repository

import (
	"sample/app/model"
	"sample/app/query"
	"sample/database"
)

type UserRepository struct {
}

func (r *UserRepository) FindAllUser(query *query.User) (users *[]model.User, err error) {
	db := database.GetDB()
	db.Preload("UserRole").Find(&users)
	return
}

func (r *UserRepository) SaveUser(u *model.User) (err error) {
	db := database.GetDB()
	err = db.Save(u).Error
	if err != nil {
		return nil
	}
	return
}
