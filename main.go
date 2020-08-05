package main

import (
	"github.com/gin-gonic/gin"
	eztemplate "github.com/michelloworld/ez-gin-template"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)

	render := eztemplate.New()
	render.TemplatesDir = "templates/"
	render.Layout = "layout/base"
	render.Debug = gin.Mode() == gin.DebugMode

	engine := gin.Default()
	engine.HTMLRender = render.Init()

	DefineRoutes(engine)
	engine.Run(":8080")

}
