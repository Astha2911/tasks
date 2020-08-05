package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type pong struct {
	Title   string
	Message string
}

func DefineRoutes(engine *gin.Engine) {
	engine.GET("/", pingHandler)
	engine.GET("/ping", pingHandler)
}

func pingHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "views/index", pong{Message: "Service is live!", Title: "welcome"})
}
