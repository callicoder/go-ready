package sqlrepo

import "github.com/callicoder/go-ready/internal/model"

type SqlUserRepository struct {
	*SqlRepository
}

func NewSqlUserRepository(sqlRepository *SqlRepository) *SqlUserRepository {
	return &SqlUserRepository{sqlRepository}
}

func (s *SqlUserRepository) Save(user *model.User) error {
	return s.DB().Save(user).Error
}

func (s *SqlUserRepository) FindById(id int) (*model.User, error) {
	var user model.User
	dbResult := s.DB().First(&user, id)
	if err := dbResult.Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *SqlUserRepository) DeleteById(id int) error {
	return s.DB().Delete(&model.User{Id: id}).Error
}
