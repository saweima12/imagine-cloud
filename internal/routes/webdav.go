package routes

import (
	"github.com/labstack/echo"
	"github.com/saweima12/imagine/internal/services"
)

func handleWebDAVRoute(handler services.WebDAVHandler) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		w, req := ctx.Response().Writer, ctx.Request()
		// pass request & response writer to webdav handler.
		handler.ServeHTTP(w, req)
		return nil
	}
}
