package handler

import (
	"github.com/labstack/echo"
	"github.com/saweima12/imagine/internal/imagine/middleware"
	"github.com/saweima12/imagine/internal/service"
)

func (r *Router) initWebDAVHandler() {
	userAuthService := service.NewUserAuthService(*r.App.UserContext)
	webDavService := service.NewWebDavService()
	// declare handler.
	handler := webDavHandler{
		WebDavService: webDavService,
	}

	// create group & add forward route.
	davGroup := r.App.Engine.Group("/webdav")
	davGroup.Use(middleware.CustomBasicAuth(userAuthService))
	davGroup.Any("/*", handler.handleWebDAVRoute)
}

type webDavHandler struct {
	WebDavService service.WebDavService
}

func (h *webDavHandler) handleWebDAVRoute(ctx echo.Context) error {
	w, req := ctx.Response().Writer, ctx.Request()
	// pass request & response writer to webdav handler.
	h.WebDavService.ServeHTTP(w, req)
	return nil
}
