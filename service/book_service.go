package service

import (
	"github.com/sandriansyafridev/golang/api/book/model/dto"
	"github.com/sandriansyafridev/golang/api/book/model/formatter"
	"github.com/sandriansyafridev/golang/api/book/model/mapping"
	"github.com/sandriansyafridev/golang/api/book/model/response"
	"github.com/sandriansyafridev/golang/api/book/repository"
)

type BookService interface {
	FindAll() ([]response.BookResponse, error)
	FindByID(BookID int) (response.BookResponse, error)
	Delete(BookID int) error
	Create(request dto.BookCreateDTO) (response.BookResponse, error)
	Update(request dto.BookUpdateDTO) (response.BookResponse, error)
}

type BookServiceImpl struct {
	repository.BookRepository
}

func (bookService *BookServiceImpl) FindAll() ([]response.BookResponse, error) {
	if books, err := bookService.BookRepository.FindAll(); err != nil {
		return nil, err
	} else {
		booksResponse := formatter.ToBooksResponse(books)
		return booksResponse, nil
	}
}

func (bookService *BookServiceImpl) FindByID(BookID int) (response.BookResponse, error) {

	if book, err := bookService.BookRepository.FindByID(BookID); err != nil {
		return response.BookResponse{}, err
	} else {
		bookResponse := formatter.ToBookResponse(book)
		return bookResponse, nil
	}
}

func (bookService *BookServiceImpl) Update(request dto.BookUpdateDTO) (response.BookResponse, error) {
	bookResponse := response.BookResponse{}
	if book, err := bookService.BookRepository.FindByID(int(request.ID)); err != nil {
		return bookResponse, err
	} else {
		book.Title = request.Title
		book.Description = request.Description
		book.Price = request.Price
		book.Rating = request.Rating

		if bookUpdated, err := bookService.BookRepository.Update(book); err != nil {
			return bookResponse, err
		} else {
			bookResponse = formatter.ToBookResponse(bookUpdated)
			return bookResponse, nil
		}
	}
}

func (bookService *BookServiceImpl) Create(request dto.BookCreateDTO) (response.BookResponse, error) {

	book := mapping.ToBookCreate(request)
	if bookCreated, err := bookService.BookRepository.Create(book); err != nil {
		return response.BookResponse{}, err
	} else {
		return formatter.ToBookResponse(bookCreated), nil
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
