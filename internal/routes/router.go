package routes

import (
	"github.com/saweima12/imagine/internal/modules"
)

type Router struct {
	App *modules.ImagineApp
}

func NewRouter(app *modules.ImagineApp) *Router {
	return &Router{
		App: app,
	}
}

func (r *Router) Init() {

	// Add static support for handle webpage.
	r.App.Engine.Static("/static", "static")

	// register webpage routes
	r.initWebRoute()

	// register api routes
	r.initApiRoute()

	// Add WebDav endpoint route & attach BasicAuth middleware.
	r.initWebDAVRoute()
}
