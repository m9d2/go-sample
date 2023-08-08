package service

import (
	"sample/app/model"
	"sample/app/query"
	"sample/app/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

func (s *UserService) FindAllUser(query *query.User) (users *[]model.User, err error) {
	return s.userRepository.FindAllUser(query)
}

func (s *UserService) SaveUser(u *model.User) (err error) {
	return s.userRepository.SaveUser(u)
}
