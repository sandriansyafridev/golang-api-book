package handler

import (
	"github.com/gin-gonic/gin"
)

type PingHandler interface {
	Ping(c *gin.Context)
}

type PingHandlerImpl struct{}

func (pingHandler *PingHandlerImpl) Ping(c *gin.Context) {
	c.JSON(200, "ping")
}

func NewPingHandler() *PingHandlerImpl {
	return &PingHandlerImpl{}
}
