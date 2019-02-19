package sqlrepo

import "github.com/callicoder/go-ready/internal/model"

type SqlUserRepository struct {
	*SqlRepository
}

func NewSqlUserRepository(sqlRepository *SqlRepository) *SqlUserRepository {
	return &SqlUserRepository{sqlRepository}
}

func (s *SqlUserRepository) Save(user *model.User) (*model.User, error) {
	if err := s.DB().Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (s *SqlUserRepository) FindById(id uint64) (*model.User, error) {
	var user model.User
	dbResult := s.DB().First(&user, id)
	if err := dbResult.Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *SqlUserRepository) DeleteById(id uint64) error {
	return s.DB().Delete(&model.User{Id: id}).Error
}
