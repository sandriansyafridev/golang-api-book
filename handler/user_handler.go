package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sandriansyafridev/golang/api/book/model/dto"
	"github.com/sandriansyafridev/golang/api/book/model/formatter"
	"github.com/sandriansyafridev/golang/api/book/service"
)

type UserHandler interface {
	FindAll(c *gin.Context)
	FindByID(c *gin.Context)
	Delete(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
}

type UserHandlerImpl struct {
	service.UserService
}

func (userHandler *UserHandlerImpl) Create(c *gin.Context) {

	request := dto.UserCreateDTO{}
	if err := c.ShouldBindJSON(&request); err != nil {
		webResponse := formatter.BuildResponseError(http.StatusNotFound, err.Error())
		c.JSON(http.StatusNotFound, webResponse)
	} else {

		if userResponse, err := userHandler.UserService.Create(request); err != nil {
			webResponse := formatter.BuildResponseError(http.StatusUnprocessableEntity, err.Error())
			c.JSON(http.StatusNotFound, webResponse)
		} else {
			userResponse.Token = "token" // next will use JWT
			webResponse := formatter.BuildResponseSuccess(http.StatusOK, userResponse)
			c.JSON(http.StatusOK, webResponse)
		}

	}
}
func (userHandler *UserHandlerImpl) Update(c *gin.Context) {

	UserID, _ := strconv.Atoi(c.Param("id"))
	request := dto.UserUpdateDTO{}
	if err := c.ShouldBindJSON(&request); err != nil {
		webResponse := formatter.BuildResponseError(http.StatusNotFound, err.Error())
		c.JSON(http.StatusNotFound, webResponse)
	} else {
		request.ID = UserID
		if userResponse, err := userHandler.UserService.Update(request); err != nil {
			webResponse := formatter.BuildResponseError(http.StatusUnprocessableEntity, err.Error())
			c.JSON(http.StatusNotFound, webResponse)
		} else {
			webResponse := formatter.BuildResponseSuccess(http.StatusOK, userResponse)
			c.JSON(http.StatusOK, webResponse)
		}
	}

}

func (userHandler *UserHandlerImpl) FindAll(c *gin.Context) {

	if usersResponse, err := userHandler.UserService.FindAll(); err != nil {
		webResponse := formatter.BuildResponseError(http.StatusNotFound, err.Error())
		c.JSON(http.StatusNotFound, webResponse)
	} else {
		webResponse := formatter.BuildResponseSuccess(http.StatusOK, usersResponse)
		c.JSON(http.StatusOK, webResponse)
	}

}
func (userHandler *UserHandlerImpl) FindByID(c *gin.Context) {
	UserID, _ := strconv.Atoi(c.Param("id"))

	if userResponse, err := userHandler.UserService.FindByID(UserID); err != nil {
		webResponse := formatter.BuildResponseError(http.StatusNotFound, err.Error())
		c.JSON(http.StatusNotFound, webResponse)
	} else {
		webResponse := formatter.BuildResponseSuccess(http.StatusOK, userResponse)
		c.JSON(http.StatusOK, webResponse)
	}

}
func (userHandler *UserHandlerImpl) Delete(c *gin.Context) {
	UserID, _ := strconv.Atoi(c.Param("id"))

	if err := userHandler.UserService.Delete(UserID); err != nil {
		webResponse := formatter.BuildResponseError(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusNotFound, webResponse)
	} else {
		webResponse := formatter.BuildResponseSuccess(http.StatusOK, gin.H{"deleted": true})
		c.JSON(http.StatusOK, webResponse)
	}

}

func NewUserHandlerImpl(userService service.UserService) *UserHandlerImpl {
	return &UserHandlerImpl{
		UserService: userService,
	}
}
