package service

import (
	"github.com/sandriansyafridev/golang/api/book/model/formatter"
	"github.com/sandriansyafridev/golang/api/book/model/response"
	"github.com/sandriansyafridev/golang/api/book/repository"
)

type BookService interface {
	FindAll() ([]response.BookResponse, error)
	FindByID(BookID int) (response.BookResponse, error)
	Delete(BookID int) error
}

type BookServiceImpl struct {
	repository.BookRepository
}

func (bookService *BookServiceImpl) FindAll() ([]response.BookResponse, error) {
	if books, err := bookService.BookRepository.FindAll(); err != nil {
		return nil, err
	} else {
		booksResponse := formatter.FormatToBooksResponse(books)
		return booksResponse, nil
	}
}

func (bookService *BookServiceImpl) FindByID(BookID int) (response.BookResponse, error) {

	if book, err := bookService.BookRepository.FindByID(BookID); err != nil {
		return response.BookResponse{}, err
	} else {
		bookResponse := formatter.FormatToBookResponse(book)
		return bookResponse, nil
	}
}

func (bookService *BookServiceImpl) Delete(BookID int) error {
	if book, err := bookService.BookRepository.FindByID(BookID); err != nil {
		return err
	} else {
		if err := bookService.BookRepository.Delete(book); err != nil {
			return err
		} else {
			return nil
		}
	}
}

func NewBookServiceImpl(bookRepository repository.BookRepository) *BookServiceImpl {
	return &BookServiceImpl{
		BookRepository: bookRepository,
	}
}
