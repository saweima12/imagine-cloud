package imagine

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/saweima12/imagine/internal/modules"
	"github.com/saweima12/imagine/internal/routes"
	"github.com/saweima12/imagine/internal/services"
)

func New() *modules.ImagineApp {
	// Initialize echo engine & services.
	engine := echo.New()
	dav := services.NewDAV()
	authService := services.NewUserAuthService()

	// Add native middleware.
	// engine.Use(middleware.Logger())

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
	app := &modules.ImagineApp{
		Engine:      engine,
		Dav:         dav,
		AuthService: authService,
	}

	// Initialize Routes
	routes.Init(app)

	return app
}
