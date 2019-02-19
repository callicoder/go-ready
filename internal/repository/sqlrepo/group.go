package sqlrepo

import "github.com/callicoder/go-ready/internal/model"

type SqlGroupRepository struct {
	*SqlRepository
}

func NewSqlGroupRepository(sqlRepository *SqlRepository) *SqlGroupRepository {
	return &SqlGroupRepository{sqlRepository}
}

func (s *SqlGroupRepository) Save(group *model.Group) (*model.Group, error) {
	if err := s.DB().Save(group).Error; err != nil {
		return nil, err
	}
	return group, nil

}

func (s *SqlGroupRepository) FindById(id uint64) (*model.Group, error) {
	var group model.Group
	dbResult := s.DB().First(&group, id)
	if err := dbResult.Error; err != nil {
		return nil, err
	}
	return &group, nil
}

func (s *SqlGroupRepository) DeleteById(id uint64) error {
	return s.DB().Delete(&model.Group{Id: id}).Error
}
