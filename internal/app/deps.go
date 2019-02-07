package app

import (
	"github.com/callicoder/go-ready/internal/config"
	"github.com/callicoder/go-ready/internal/repository"
	"github.com/callicoder/go-ready/internal/repository/sqlrepo"
	"github.com/callicoder/go-ready/internal/service"
)

type Dependencies struct {
	Repository     repository.Repository
	UserRepository repository.UserRepository
	UserService    service.UserService
	GroupService   service.GroupService
	TokenService   service.TokenService
}

func NewDependencies(config *config.Config) (*Dependencies, error) {
	sqlRepository, err := sqlrepo.New(config.Database)
	if err != nil {
		return nil, err
	}
	userRepository := sqlrepo.NewSqlUserRepository(sqlRepository)
	groupRepository := sqlrepo.NewSqlGroupRepository(sqlRepository)

	userService := service.NewUserService(userRepository)
	groupService := service.NewGroupService(groupRepository)
	tokenService := service.NewTokenService(config.Auth)

	return &Dependencies{
		Repository:     sqlRepository,
		UserRepository: userRepository,
		UserService:    userService,
		TokenService:   tokenService,
		GroupService:   groupService,
	}, nil
}
