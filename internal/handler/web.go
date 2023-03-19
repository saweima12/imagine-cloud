package handler

import (
	"github.com/gin-gonic/contrib/static"
)

func (r *Router) initWebHandler() {

	engine := r.App.Engine

	engine.Use(static.Serve("/", static.LocalFile("web", true)))

	engine.Static("/assets", "/web/assets")
}
