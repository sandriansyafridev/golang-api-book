package repository

import (
	"github.com/sandriansyafridev/golang/api/book/model/entity"
	"gorm.io/gorm"
)

type BookRepository interface {
	FindAll() ([]entity.Book, error)
	FindByID(BookID int) (entity.Book, error)
	Delete(book entity.Book) error
	Create(book entity.Book) (entity.Book, error)
	Update(book entity.Book) (entity.Book, error)
}

type BookRepositoryImpl struct {
	DB *gorm.DB
}

func (bookRepository *BookRepositoryImpl) FindAll() ([]entity.Book, error) {

	books := []entity.Book{}
	if err := bookRepository.DB.Find(&books).Error; err != nil {
		return books, err
	} else {
		return books, nil
	}
}

func (bookRepository *BookRepositoryImpl) FindByID(BookID int) (entity.Book, error) {

	book := entity.Book{}
	if err := bookRepository.DB.First(&book, BookID).Error; err != nil {
		return book, err
	} else {
		return book, nil
	}
}

func (bookRepository *BookRepositoryImpl) Delete(book entity.Book) error {
	if err := bookRepository.DB.Delete(&book).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func (bookRepository *BookRepositoryImpl) Create(book entity.Book) (entity.Book, error) {

	if err := bookRepository.DB.Create(&book).Error; err != nil {
		return book, err
	} else {
		return book, nil
	}

}

func (bookRepository *BookRepositoryImpl) Update(book entity.Book) (entity.Book, error) {

	if err := bookRepository.DB.Save(&book).Error; err != nil {
		return book, err
	} else {
		return book, nil
	}

}

func NewBookRepositoryImpl(db *gorm.DB) *BookRepositoryImpl {
	return &BookRepositoryImpl{
		DB: db,
	}
}
