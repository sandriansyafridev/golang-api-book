package main

import (
	"github.com/sandriansyafridev/golang/api/book/app"
	"github.com/sandriansyafridev/golang/api/book/handler"
)

var (
	db              = app.NewDatabase()
	pingHandlerImpl = handler.NewPingHandler()
)

func main() {
	defer app.CloseDatabase(db)

	r := app.NewRouter(pingHandlerImpl)
	r.Run(":8080")
}
