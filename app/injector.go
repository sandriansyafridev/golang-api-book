//go:build wireinject
//+build wireinject

package app

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/sandriansyafridev/golang/api/book/handler"
	"github.com/sandriansyafridev/golang/api/book/repository"
	"github.com/sandriansyafridev/golang/api/book/service"
)

var (
	BookSet = wire.NewSet(
		wire.Bind(new(repository.BookRepository), new(*repository.BookRepositoryImpl)),
		wire.Bind(new(service.BookService), new(*service.BookServiceImpl)),
		wire.Bind(new(handler.BookHandler), new(*handler.BookHandlerImpl)),
		repository.NewBookRepositoryImpl,
		service.NewBookServiceImpl,
		handler.NewBookHandlerImpl,
	)
)

func InitializedRouter() *gin.Engine {

	wire.Build(
		NewDatabase,
		NewRouter,
		BookSet,
	)

	return nil
}
