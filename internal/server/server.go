package server

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/saweima12/imagine/internal/handler"
	"github.com/saweima12/imagine/internal/imagine"
	"github.com/saweima12/imagine/internal/imagine/config"
	"github.com/saweima12/imagine/internal/service"
)

func New() *imagine.ServerApp {
	// Initialize echo engine & service.
	engine := gin.Default()
	dav := service.NewWebDavService()

	// Initialize config
	userContext := config.LoadFromEnv()

	// Add native middleware.

	// Add CORS middleware & add support methods
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.AllowMethods = append(corsConfig.AllowMethods,
		"LOCK",
		"UNLOCK",
		"COPY",
		"MOVE",
		"MKCOL",
		"PROPFIND",
		"PROPPATCH",
	)
	fmt.Println(corsConfig)
	engine.Use(cors.New(corsConfig))

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
