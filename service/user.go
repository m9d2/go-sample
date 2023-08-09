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

func (s *UserService) FindAll(req *vo.FindAllUserReq) (users *[]vo.FindAllUserRsp, err error) {
	db := database.GetDB()
	return s.userRepository.FindAllUser(db, req)
}

func (s *UserService) Save(vo *vo.SaveUserReq) error {
	db := database.GetDB()
	err := db.Transaction(func(tx *gorm.DB) error {
		user := model.User{
			Name:   vo.Name,
			DeptId: vo.DeptId,
		}
		u, err := s.userRepository.SaveUser(tx, &user)
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
