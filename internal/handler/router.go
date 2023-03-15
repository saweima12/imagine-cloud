package handler

import "github.com/saweima12/imagine/internal/imagine"

type Router struct {
	App *imagine.ServerApp
}

func NewRouter(app *imagine.ServerApp) *Router {
	return &Router{
		App: app,
	}
}

func (r *Router) Init() {

	// Add static support for handle webpage.
	r.App.Engine.Static("/static", "static")

	// register webpage routes
	r.initWebHandler()

	// register api routes
	apiGroup := r.App.Engine.Group("/api/v1")
	r.initApiHandler(apiGroup)

	// Add WebDav endpoint route & attach BasicAuth middleware.
	r.initWebDAVHandler()
}
