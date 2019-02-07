package service

import "github.com/callicoder/go-ready/internal/repository"

type GroupService interface {
}

type groupService struct {
	groupRepository repository.GroupRepository
}

func NewGroupService(groupRepository repository.GroupRepository) GroupService {
	return &groupService{
		groupRepository: groupRepository,
	}
}
