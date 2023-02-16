package routes

import (
	echoMI "github.com/labstack/echo/middleware"
)

func (r *Router) initWebRoute() {

	webGroup := r.App.Engine.Group("/")
	webGroup.Use(echoMI.StaticWithConfig(echoMI.StaticConfig{
		Root:   "web",
		Browse: false,
		HTML5:  true,
	}))

}
