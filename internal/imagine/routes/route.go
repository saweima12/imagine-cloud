package routes

import (
	"github.com/labstack/echo"
	"github.com/saweima12/imagine/internal/imagine"
)

func Register(s *imagine.Server) {
	// register webdav handler.
	s.Echo.Any("/webdav", func(ctx echo.Context) error {
		s.Dav.ServeHTTP(&ctx.Response().Writer, ctx.Request())
		return nil
	})
}
