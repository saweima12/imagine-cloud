package routes

import (
	"github.com/labstack/echo"
	echoMI "github.com/labstack/echo/middleware"
	"github.com/saweima12/imagine/internal/modules"
	"github.com/saweima12/imagine/internal/modules/middleware"
)

func Init(app *modules.ImagineApp) {

	// Add static support for handle webpage.
	app.Engine.Static("/static", "static")

	webGroup := app.Engine.Group("/")
	webGroup.Use(echoMI.StaticWithConfig(echoMI.StaticConfig{
		Root:   "web",
		Browse: false,
		HTML5:  true,
	}))

	// Add Login & Operate Endpoint.
	apiGroup := app.Engine.Group("/api/v1")
	apiGroup.POST("/login", userlogin(app))

	operateGroup := apiGroup.Group("")
	operateGroup.Use(middleware.CustomBasicAuth(app))

	operateGroup.GET("/appa", func(ctx echo.Context) error {
		return nil
	})

	// Add WebDav route & attach BasicAuth middleware.
	davGroup := app.Engine.Group("/webdav")
	davGroup.Use(middleware.CustomBasicAuth(app))
	davGroup.Any("/*", handleWebDAVRoute(app.Dav))

}
