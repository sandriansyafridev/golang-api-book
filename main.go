package main

import "github.com/sandriansyafridev/golang/api/book/app"

func main() {

	r := app.InitializedRouter()
	r.Run(":8080")

}
