package repository

import (
	"gorm.io/gorm"
	"sample/model"
)

type UserRoleRepository struct {
}

func (r *UserRoleRepository) Save(db *gorm.DB, role *model.UserRole) error {
	return db.Create(&role).Error
}

func (r *UserRoleRepository) Delete(db *gorm.DB, id uint) error {
	return db.Delete(id).Error
}

func (r *UserRoleRepository) DeleteByUserId(db *gorm.DB, userId uint) error {
	return db.Where("user_id = ?", userId).Delete(&model.UserRole{}).Error
}
