package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sandriansyafridev/golang/api/book/handler"
)

func NewRouter(bookHandler handler.BookHandler, userHandler handler.UserHandler) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	v1.GET("/books", bookHandler.FindAll)
	v1.POST("/books", bookHandler.Create)
	v1.GET("/books/:id", bookHandler.FindByID)
	v1.PUT("/books/:id", bookHandler.Update)
	v1.DELETE("/books/:id", bookHandler.Delete)

	v1.GET("/users", userHandler.FindAll)
	v1.GET("/users/:id", userHandler.FindByID)
	v1.DELETE("/users/:id", userHandler.Delete)

	return r
}
