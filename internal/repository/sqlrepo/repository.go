package sqlrepo

import (
	"github.com/callicoder/go-ready/internal/config"
	"github.com/callicoder/go-ready/internal/repository"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type SqlRepository struct {
	db *gorm.DB
}

func New(dbConfig config.DatabaseConfig) (*SqlRepository, error) {
	db, err := gorm.Open(dbConfig.Driver, dbConfig.URL())
	if err != nil {
		return nil, err
	}

	db.DB().SetMaxIdleConns(0)

	repository := newRepository(db)
	return repository, nil
}

func newRepository(db *gorm.DB) *SqlRepository {
	repository := &SqlRepository{db: db}
	return repository
}

func (s *SqlRepository) DB() *gorm.DB {
	return s.db
}

func (s *SqlRepository) Begin() (repository.Repository, error) {
	tx := s.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	repository := newRepository(tx)
	return repository, nil
}

func (s *SqlRepository) Commit() error {
	c := s.db.Commit()
	if c.Error != nil {
		return c.Error
	}
	return nil
}

func (s *SqlRepository) Rollback() error {
	c := s.db.Rollback()
	if c.Error != nil {
		return c.Error
	}
	return nil
}

func (s *SqlRepository) Close() {
	s.db.Close()
}
