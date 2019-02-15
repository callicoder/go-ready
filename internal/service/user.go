package service

import (
	"github.com/callicoder/go-ready/internal/model"
	"github.com/callicoder/go-ready/internal/repository"
)

type UserService interface {
	Save(user *model.User) (*model.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) Save(user *model.User) (*model.User, error) {
	return nil, nil
}
