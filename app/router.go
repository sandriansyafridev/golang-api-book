package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sandriansyafridev/golang/api/book/handler"
)

func NewRouter(bookHandler handler.BookHandler, userHandler handler.UserHandler, memberHandler handler.MemberHandler) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	v1.GET("/books", bookHandler.FindAll)
	v1.POST("/books", bookHandler.Create)
	v1.GET("/books/:id", bookHandler.FindByID)
	v1.PUT("/books/:id", bookHandler.Update)
	v1.DELETE("/books/:id", bookHandler.Delete)

	v1.GET("/users", userHandler.FindAll)
	v1.POST("/users", userHandler.Create)
	v1.GET("/users/:id", userHandler.FindByID)
	v1.PUT("/users/:id", userHandler.Update)
	v1.DELETE("/users/:id", userHandler.Delete)

	v1.GET("/members", memberHandler.FindAll)
	v1.POST("/members", memberHandler.Create)
	v1.GET("/members/:id", memberHandler.FindByID)
	v1.PUT("/members/:id", memberHandler.Update)
	v1.DELETE("/members/:id", memberHandler.Delete)

	return r
}
