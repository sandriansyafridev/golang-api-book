package main

import (
	"github.com/sandriansyafridev/golang/api/book/app"
	"github.com/sandriansyafridev/golang/api/book/handler"
	"github.com/sandriansyafridev/golang/api/book/repository"
	"github.com/sandriansyafridev/golang/api/book/service"
)

var (
	db             = app.NewDatabase()
	bookRepository = repository.NewBookRepositoryImpl(db)
	bookService    = service.NewBookServiceImpl(bookRepository)
	bookHandler    = handler.NewBookHandlerImpl(bookService)
)

func main() {
	defer app.CloseDatabase(db)

	r := app.NewRouter(bookHandler)
	r.Run(":8080")
}
