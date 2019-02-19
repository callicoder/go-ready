package service

import (
	"github.com/callicoder/go-ready/internal/model"
	"github.com/callicoder/go-ready/internal/repository"
)

type GroupService interface {
	Save(group *model.Group) (*model.Group, error)
}

type groupService struct {
	groupRepository repository.GroupRepository
}

func NewGroupService(groupRepository repository.GroupRepository) GroupService {
	return &groupService{
		groupRepository: groupRepository,
	}
}

func (s *groupService) Save(group *model.Group) (*model.Group, error) {
	return s.groupRepository.Save(group)
}
