package repository

import (
	"github.com/callicoder/go-ready/internal/model"
)

type Repository interface {
	Begin() (Repository, error)
	Commit() error
	Rollback() error
	Close()
}

type UserRepository interface {
	Save(user *model.User) (*model.User, error)
	FindById(id uint64) (*model.User, error)
	DeleteById(id uint64) error
}

type GroupRepository interface {
	Save(group *model.Group) (*model.Group, error)
	FindById(id uint64) (*model.Group, error)
	DeleteById(id uint64) error
}
