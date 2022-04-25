package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sandriansyafridev/golang/api/book/model/dto"
	"github.com/sandriansyafridev/golang/api/book/model/formatter"
	"github.com/sandriansyafridev/golang/api/book/service"
)

type BookHandler interface {
	FindAll(c *gin.Context)
	FindByID(c *gin.Context)
	Delete(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
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
func (bookHandler *BookHandlerImpl) Create(c *gin.Context) {

	request := dto.BookCreateDTO{}
	if err := c.ShouldBindJSON(&request); err != nil {
		webResponse := formatter.BuildResponseError(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusNotFound, webResponse)
		return
	}

	if bookResponse, err := bookHandler.BookService.Create(request); err != nil {
		webResponse := formatter.BuildResponseError(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusNotFound, webResponse)
	} else {
		webResponse := formatter.BuildResponseSuccess(http.StatusOK, bookResponse)
		c.JSON(http.StatusOK, webResponse)
	}
}

func (bookHandler *BookHandlerImpl) Update(c *gin.Context) {
	BookID, _ := strconv.Atoi(c.Param("id"))
	request := dto.BookUpdateDTO{}
	if err := c.ShouldBindJSON(&request); err != nil {
		webResponse := formatter.BuildResponseError(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusNotFound, webResponse)
		return
	}

	request.ID = uint64(BookID)

	if bookResponse, err := bookHandler.BookService.Update(request); err != nil {
		webResponse := formatter.BuildResponseError(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusNotFound, webResponse)
	} else {
		webResponse := formatter.BuildResponseSuccess(http.StatusOK, bookResponse)
		c.JSON(http.StatusOK, webResponse)
	}

}

func (bookHandler *BookHandlerImpl) Delete(c *gin.Context) {
	BookID, _ := strconv.Atoi(c.Param("id"))

	if err := bookHandler.BookService.Delete(BookID); err != nil {
		webResponse := formatter.BuildResponseError(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusNotFound, webResponse)
	} else {
		webResponse := formatter.BuildResponseSuccess(http.StatusOK, gin.H{"deleted": true})
		c.JSON(http.StatusOK, webResponse)
	}

}

func NewBookHandlerImpl(bookService service.BookService) *BookHandlerImpl {
	return &BookHandlerImpl{
		BookService: bookService,
	}
}
