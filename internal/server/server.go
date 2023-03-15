package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/saweima12/imagine/internal/handler"
	"github.com/saweima12/imagine/internal/imagine"
	"github.com/saweima12/imagine/internal/imagine/config"
	"github.com/saweima12/imagine/internal/service"
)

func New() *imagine.ServerApp {
	// Initialize echo engine & service.
	engine := echo.New()
	dav := service.NewWebDavService()

	// Initialize config
	userContext := config.LoadFromEnv()

	// Add native middleware.
	engine.Use(middleware.Logger())

	// Add CORS middleware & add support methods
	defaultMethods := middleware.DefaultCORSConfig.AllowMethods
	extMethods := append(defaultMethods, "LOCK", "MOVE", "PROPFIND", "UNLOCK", "PROPPATCH", "MKCOL", "LOCK")

	engine.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowHeaders: middleware.DefaultCORSConfig.AllowHeaders,
		AllowMethods: extMethods,
	}))

	// Create App instance.
	app := &imagine.ServerApp{
		UserContext: userContext,
		Engine:      engine,
		Dav:         dav,
	}

	// Initialize Routes
	router := handler.NewRouter(app)
	router.Init()

	return app
}
