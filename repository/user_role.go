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
