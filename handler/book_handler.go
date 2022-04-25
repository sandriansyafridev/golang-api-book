package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sandriansyafridev/golang/api/book/model/formatter"
	"github.com/sandriansyafridev/golang/api/book/service"
)

type BookHandler interface {
	FindAll(c *gin.Context)
	FindByID(c *gin.Context)
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

func (bookHandler *BookHandlerImpl) FindByID(c *gin.Context) {
	BookID, _ := strconv.Atoi(c.Param("id"))

	if bookResponse, err := bookHandler.BookService.FindByID(BookID); err != nil {
		webResponse := formatter.BuildResponseError(http.StatusNotFound, err.Error())
		c.JSON(http.StatusNotFound, webResponse)
	} else {
		webResponse := formatter.BuildResponseSuccess(http.StatusOK, bookResponse)
		c.JSON(http.StatusOK, webResponse)
	}

}

func NewBookHandlerImpl(bookService service.BookService) *BookHandlerImpl {
	return &BookHandlerImpl{
		BookService: bookService,
	}
}
