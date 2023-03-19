package handler

import (
	"log"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/saweima12/imagine/internal/imagine/middleware"
	"github.com/saweima12/imagine/internal/service"
)

func (r *Router) initWebDAVHandler() {
	userAuthService := service.NewUserAuthService(*r.App.UserContext)
	webDavService := service.NewWebDavService()
	// declare regex pattern.
	pattern := regexp.MustCompile(`^/webdav`)

	// declare handler.
	handler := webDavHandler{
		WebDavService: webDavService,
		Pattern:       pattern,
	}
	// declare extension Method.
	extMethodList := []string{
		"PROPFIND", "PROPPATCH", "LOCK", "UNLOCK", "MOVE", "MKCOL",
	}
	// create group & add forwarding route.
	davGroup := r.App.Engine.Group("/webdav")
	davGroup.Use(middleware.CustomBasicAuth(userAuthService))
	{
		davGroup.Any("/*any", handler.handleWebDAVRoute)
		for _, method := range extMethodList {
			davGroup.Handle(method, "/*any", handler.handleWebDAVRoute)
		}
	}
}

type webDavHandler struct {
	WebDavService service.WebDavService

	Pattern *regexp.Regexp
}

func (h *webDavHandler) handleWebDAVRoute(ctx *gin.Context) {
	w, req := ctx.Writer, ctx.Request
	log.Printf("Path: %s", req.URL.Path)
	// remove url prefix.
	req.URL.Path = h.Pattern.ReplaceAllString(req.URL.Path, "")

	// pass request & response writer to webdav handler.
	log.Printf("=================")
	log.Printf("Path: %s", req.URL.Path)
	log.Printf("Method: %s", req.Method)
	log.Printf("Cookie: %s", ctx.Request.Cookies())
	log.Printf("=================")

	h.WebDavService.ServeHTTP(w, req)

}
