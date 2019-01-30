package service

import "github.com/callicoder/go-ready/internal/repository"

type GroupService struct {
	groupRepository repository.GroupRepository
}

func NewGroupService(groupRepository repository.GroupRepository) *GroupService {
	return &GroupService{
		groupRepository: groupRepository,
	}
}
