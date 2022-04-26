// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package app

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/sandriansyafridev/golang/api/book/handler"
	"github.com/sandriansyafridev/golang/api/book/repository"
	"github.com/sandriansyafridev/golang/api/book/service"
)

// Injectors from injector.go:

func InitializedRouter() *gin.Engine {
	db := NewDatabase()
	bookRepositoryImpl := repository.NewBookRepositoryImpl(db)
	bookServiceImpl := service.NewBookServiceImpl(bookRepositoryImpl)
	bookHandlerImpl := handler.NewBookHandlerImpl(bookServiceImpl)
	engine := NewRouter(bookHandlerImpl)
	return engine
}

// injector.go:

var (
	BookSet = wire.NewSet(wire.Bind(new(repository.BookRepository), new(*repository.BookRepositoryImpl)), wire.Bind(new(service.BookService), new(*service.BookServiceImpl)), wire.Bind(new(handler.BookHandler), new(*handler.BookHandlerImpl)), repository.NewBookRepositoryImpl, service.NewBookServiceImpl, handler.NewBookHandlerImpl)
)
