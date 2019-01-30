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
	Save(user *model.User) error
	FindById(id int) (*model.User, error)
	DeleteById(id int) error
}

type GroupRepository interface {
	Save(group *model.Group) error
	FindById(id int) (*model.Group, error)
	DeleteById(id int) error
}
