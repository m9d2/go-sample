package service

import (
	"gorm.io/gorm"
	"sample/database"
	"sample/model"
	"sample/model/vo"
	"sample/repository"
)

type UserService struct {
	userRepository     repository.UserRepository
	UserRoleRepository repository.UserRoleRepository
}

func (s *UserService) FindById(id uint) (user *vo.FindAllUserRsp, err error) {
	db := database.GetDB()
	return s.userRepository.FindById(db, id)
}

func (s *UserService) FindAll(req *vo.FindAllUserReq) (users *[]vo.FindAllUserRsp, err error) {
	db := database.GetDB()
	return s.userRepository.FindAll(db, req)
}

func (s *UserService) Save(vo *vo.SaveUserReq) error {
	db := database.GetDB()
	err := db.Transaction(func(tx *gorm.DB) error {
		user := model.User{
			Name:   vo.Name,
			DeptId: vo.DeptId,
		}
		u, err := s.userRepository.Save(tx, &user)
		if err != nil {
			return err
		}

		userRole := model.UserRole{
			UserID: u.ID,
			RoleID: vo.RoleId,
		}
		err = s.UserRoleRepository.Save(tx, &userRole)
		if err != nil {
			return err
		}
		return err
	})
	if err != nil {
		return err
	}
	return err
}

func (s *UserService) Update(vo *vo.SaveUserReq) error {
	db := database.GetDB()
	err := db.Transaction(func(tx *gorm.DB) error {
		user := model.User{
			Model: model.Model{
				ID: vo.ID,
			},
			Name:   vo.Name,
			DeptId: vo.DeptId,
		}
		_, err := s.userRepository.Update(tx, &user)
		if err != nil {
			return err
		}

		err = s.UserRoleRepository.DeleteByUserId(tx, vo.ID)
		if err != nil {
			return err
		}

		userRole := model.UserRole{
			UserID: vo.ID,
			RoleID: vo.RoleId,
		}

		err = s.UserRoleRepository.Save(tx, &userRole)
		if err != nil {
			return err
		}
		return err
	})
	if err != nil {
		return err
	}
	return err
}

func (s UserService) Delete(id uint) error {
	db := database.GetDB()
	err := db.Transaction(func(tx *gorm.DB) error {
		err := s.userRepository.Delete(tx, id)
		if err != nil {
			return err
		}

		err = s.UserRoleRepository.DeleteByUserId(tx, id)
		if err != nil {
			return err
		}
		return err
	})
	return err
}
