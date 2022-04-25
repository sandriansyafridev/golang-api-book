package repository

import (
	"github.com/sandriansyafridev/golang/api/book/model/entity"
	"gorm.io/gorm"
)

type BookRepository interface {
	FindAll() ([]entity.Book, error)
}

type BookRepositoryImpl struct {
	DB *gorm.DB
}

func (bookRepository *BookRepositoryImpl) FindAll() ([]entity.Book, error) {

	books := []entity.Book{}
	if err := bookRepository.DB.Find(&books).Error; err != nil {
		return books, err
	}

	return books, nil
}

func NewBookRepositoryImpl(db *gorm.DB) *BookRepositoryImpl {
	return &BookRepositoryImpl{
		DB: db,
	}
}
