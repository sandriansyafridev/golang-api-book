package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sandriansyafridev/golang/api/book/handler"
)

func NewRouter(pingHandler handler.PingHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", pingHandler.Ping)

	return r
}
