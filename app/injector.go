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

	UserSet = wire.NewSet(
		wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
		wire.Bind(new(service.UserService), new(*service.UserServiceImpl)),
		wire.Bind(new(handler.UserHandler), new(*handler.UserHandlerImpl)),
		repository.NewUserRepositoryImpl,
		service.NewUserServiceImpl,
		handler.NewUserHandlerImpl,
	)

	MemberSet = wire.NewSet(
		wire.Bind(new(repository.MemberRepository), new(*repository.MemberRepositoryImpl)),
		wire.Bind(new(service.MemberService), new(*service.MemberServiceImpl)),
		wire.Bind(new(handler.MemberHandler), new(*handler.MemberHandlerImpl)),
		repository.NewMemberRepositoryImpl,
		service.NewMemberServiceImpl,
		handler.NewMemberHandlerImpl,
	)
)

func InitializedRouter() *gin.Engine {

	wire.Build(
		NewDatabase,
		NewRouter,
		BookSet,
		UserSet,
		MemberSet,
	)

	return nil
}
