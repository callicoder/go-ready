package service

import (
	"github.com/callicoder/go-ready/internal/model"
	"github.com/callicoder/go-ready/internal/repository"
)

type UserService interface {
	Create(user *model.User) (*model.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) Create(user *model.User) (*model.User, error) {
	return nil, nil
}
