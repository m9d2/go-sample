package repository

import (
	"gorm.io/gorm"
	"sample/model"
	"sample/model/vo"
)

type UserRepository struct {
}

func (r *UserRepository) FindAll(db *gorm.DB, req *vo.FindAllUserReq) (users *[]vo.FindAllUserRsp, err error) {
	tx := db.Table(model.TABLE_NAME_USER).
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

func (r *UserRepository) Save(db *gorm.DB, m *model.User) (user *model.User, err error) {
	err = db.Create(&m).Error
	return m, err
}

func (r *UserRepository) Update(db *gorm.DB, m *model.User) (user *model.User, err error) {
	err = db.Save(&m).Error
	return
}

func (r UserRepository) Delete(db *gorm.DB, id uint) error {
	return db.Where("id = ?", id).Delete(&model.User{}).Error
}

func (r *UserRepository) FindById(db *gorm.DB, id uint) (req *vo.FindAllUserRsp, err error) {
	err = db.Table(model.TABLE_NAME_USER).Where("id = ?", id).Find(&req).Error
	return
}
