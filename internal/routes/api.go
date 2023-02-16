package routes

import (
	"github.com/labstack/echo"
	"github.com/saweima12/imagine/internal/modules"
	"github.com/saweima12/imagine/internal/modules/config"
	"github.com/saweima12/imagine/internal/modules/middleware"
)

func (r *Router) initApiRoute() {

	app := r.App
	apiGroup := app.Engine.Group("/api/v1")
	// Add Login & Operate Endpoint.
	apiGroup.POST("/login", userlogin(app))

	operateGroup := apiGroup.Group("")
	operateGroup.Use(middleware.CustomBasicAuth(app))
}

func userlogin(app *modules.ImagineApp) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// Bind userContext to model
		user := new(config.UserContext)
		if err := ctx.Bind(user); err != nil {
			return ctx.String(400, "Bad Request.")
		}

		token, err := app.AuthService.VerifyUser(user.Username, user.Password)

		if err != nil {
			ctx.String(400, "Username or password Error.")
			return err
		}

		ctx.JSON(200, modules.LoginResponse{Token: token})

		return nil
	}
}
