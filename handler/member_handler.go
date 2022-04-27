package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sandriansyafridev/golang/api/book/model/dto"
	"github.com/sandriansyafridev/golang/api/book/model/formatter"
	"github.com/sandriansyafridev/golang/api/book/service"
)

type MemberHandler interface {
	FindAll(c *gin.Context)
	FindByID(c *gin.Context)
	Delete(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
}

type MemberHandlerImpl struct {
	service.MemberService
}

func (memberHandler *MemberHandlerImpl) FindAll(c *gin.Context) {
	if members, err := memberHandler.MemberService.FindAll(); err != nil {
		webResponse := formatter.BuildResponseError(http.StatusNotFound, err.Error())
		c.JSON(http.StatusNotFound, webResponse)
	} else {
		webResponse := formatter.BuildResponseSuccess(http.StatusOK, members)
		c.JSON(http.StatusOK, webResponse)
	}
}

func (memberHandler *MemberHandlerImpl) FindByID(c *gin.Context) {

	MemberID, _ := strconv.Atoi(c.Param("id"))
	if memberResponse, err := memberHandler.MemberService.FindByID(MemberID); err != nil {
		webResponse := formatter.BuildResponseError(http.StatusNotFound, err.Error())
		c.JSON(http.StatusNotFound, webResponse)
	} else {
		webResponse := formatter.BuildResponseSuccess(http.StatusOK, memberResponse)
		c.JSON(http.StatusOK, webResponse)
	}

}

func (memberHandler *MemberHandlerImpl) Delete(c *gin.Context) {
	MemberID, _ := strconv.Atoi(c.Param("id"))

	if err := memberHandler.MemberService.Delete(MemberID); err != nil {
		webResponse := formatter.BuildResponseError(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusNotFound, webResponse)
	} else {
		webResponse := formatter.BuildResponseSuccess(http.StatusOK, gin.H{"deleted": true})
		c.JSON(http.StatusOK, webResponse)
	}

}
func (memberHandler *MemberHandlerImpl) Create(c *gin.Context) {
	request := dto.MemberCreateDTO{}
	if err := c.ShouldBindJSON(&request); err != nil {
		webResponse := formatter.BuildResponseError(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusNotFound, webResponse)
		return
	}

	if memberResponse, err := memberHandler.MemberService.Create(request); err != nil {
		webResponse := formatter.BuildResponseError(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusNotFound, webResponse)
	} else {
		webResponse := formatter.BuildResponseSuccess(http.StatusOK, memberResponse)
		c.JSON(http.StatusOK, webResponse)
	}
}
func (memberHandler *MemberHandlerImpl) Update(c *gin.Context) {

	MemberID, _ := strconv.Atoi(c.Param("id"))
	request := dto.MemberUpdateDTO{}
	if err := c.ShouldBindJSON(&request); err != nil {
		webResponse := formatter.BuildResponseError(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusNotFound, webResponse)
		return
	}

	request.ID = MemberID
	if memberResponse, err := memberHandler.MemberService.Update(request); err != nil {
		webResponse := formatter.BuildResponseError(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusNotFound, webResponse)
	} else {
		webResponse := formatter.BuildResponseSuccess(http.StatusOK, memberResponse)
		c.JSON(http.StatusOK, webResponse)
	}

}

func NewMemberHandlerImpl(memberService service.MemberService) *MemberHandlerImpl {
	return &MemberHandlerImpl{
		MemberService: memberService,
	}
}
