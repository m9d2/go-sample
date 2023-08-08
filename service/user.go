package service

import (
	"sample/model"
	"sample/model/vo"
	"sample/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

func (s *UserService) FindAllUser(query *vo.User) (users *[]model.User, err error) {
	return s.userRepository.FindAllUser(query)
}

func (s *UserService) SaveUser(u *model.User) (err error) {
	return s.userRepository.SaveUser(u)
}
