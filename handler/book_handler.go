package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sandriansyafridev/golang/api/book/model/formatter"
	"github.com/sandriansyafridev/golang/api/book/service"
)

type BookHandler interface {
	FindAll(c *gin.Context)
}

type BookHandlerImpl struct {
	service.BookService
}

func (bookHandler *BookHandlerImpl) FindAll(c *gin.Context) {

	if booksResponse, err := bookHandler.BookService.FindAll(); err != nil {
		webResponse := formatter.BuildResponseError(http.StatusNotFound, err.Error())
		c.JSON(http.StatusNotFound, webResponse)
	} else {
		webResponse := formatter.BuildResponseSuccess(http.StatusOK, booksResponse)
		c.JSON(http.StatusOK, webResponse)
	}

}

func NewBookHandlerImpl(bookService service.BookService) *BookHandlerImpl {
	return &BookHandlerImpl{
		BookService: bookService,
	}
}
