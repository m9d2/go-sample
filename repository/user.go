package repository

import (
	"gorm.io/gorm"
	"sample/model"
	"sample/model/vo"
)

type UserRepository struct {
}

func (r *UserRepository) FindAllUser(db *gorm.DB, req *vo.FindAllUserReq) (users *[]vo.FindAllUserRsp, err error) {
	tx := db.Table("user").
		Select("user.*, role.name as roleName").
		Joins("left join user_role on user.id = user_role.user_id").
		Joins("left join role on role.id = user_role.role_id").
		Offset(req.Page).
		Limit(req.Size)

	if req.Name != "" {
		tx.Where("user.name = ?", req.Name)
	}
	err = tx.Find(&users).Error
	return
}

func (r *UserRepository) SaveUser(db *gorm.DB, m *model.User) (user *model.User, err error) {
	err = db.Create(&m).Error
	return m, err
}
