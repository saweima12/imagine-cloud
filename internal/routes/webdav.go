package routes

import (
	"github.com/labstack/echo"
	"github.com/saweima12/imagine/internal/modules/middleware"
	"github.com/saweima12/imagine/internal/services"
)

func (r *Router) initWebDAVRoute() {

	app := r.App
	// create group & add forward route.
	davGroup := r.App.Engine.Group("/webdav")
	davGroup.Use(middleware.CustomBasicAuth(app))
	davGroup.Any("/*", handleWebDAVRoute(app.Dav))
}

func handleWebDAVRoute(handler services.WebDAVHandler) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		w, req := ctx.Response().Writer, ctx.Request()
		// pass request & response writer to webdav handler.
		handler.ServeHTTP(w, req)
		return nil
	}
}
