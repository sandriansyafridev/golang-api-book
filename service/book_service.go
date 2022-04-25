package service

import (
	"github.com/sandriansyafridev/golang/api/book/model/formatter"
	"github.com/sandriansyafridev/golang/api/book/model/response"
	"github.com/sandriansyafridev/golang/api/book/repository"
)

type BookService interface {
	FindAll() ([]response.BookResponse, error)
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

func NewBookServiceImpl(bookRepository repository.BookRepository) *BookServiceImpl {
	return &BookServiceImpl{
		BookRepository: bookRepository,
	}
}
