package mapping

import (
	"github.com/sandriansyafridev/golang/api/book/model/dto"
	"github.com/sandriansyafridev/golang/api/book/model/entity"
)

func ToBook(request dto.BookCreateDTO) entity.Book {
	return entity.Book{
		Title:       request.Title,
		Description: request.Description,
		Price:       request.Price,
		Rating:      request.Rating,
	}
}
