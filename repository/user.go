package repository

import (
	"sample/database"
	"sample/model"
	"sample/model/vo"
)

type UserRepository struct {
}

func (r *UserRepository) FindAllUser(vo *vo.User) (users *[]model.User, err error) {
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
