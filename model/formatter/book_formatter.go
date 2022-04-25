package formatter

import (
	"github.com/sandriansyafridev/golang/api/book/model/entity"
	"github.com/sandriansyafridev/golang/api/book/model/response"
)

func FormatToBookResponse(book entity.Book) response.BookResponse {

	return response.BookResponse{
		ID:          book.ID,
		Description: book.Description,
		Title:       book.Title,
		Price:       book.Price,
		Rating:      book.Rating,
		CreatedAt:   book.CreatedAt,
		UpdatedAt:   book.UpdatedAt,
	}
}

func FormatToBooksResponse(books []entity.Book) []response.BookResponse {

	booksResponse := []response.BookResponse{}
	for _, book := range books {
		bookResponse := FormatToBookResponse(book)
		booksResponse = append(booksResponse, bookResponse)
	}

	return booksResponse
}
