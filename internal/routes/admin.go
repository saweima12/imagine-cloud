package routes

import (
	"github.com/labstack/echo"
	"github.com/saweima12/imagine/internal/modules"
	"github.com/saweima12/imagine/internal/modules/config"
)

func userlogin(app *modules.ImagineApp) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// Bind userContext to model
		user := new(config.UserContext)
		if err := ctx.Bind(user); err != nil {
			return ctx.String(400, "Bad Request.")
		}

		ok, err := app.AuthService.VerifyUser(user.Username, user.Password)

		if !ok {
			ctx.String(400, "Password Error.")
			return nil
		}

		if err != nil {
			panic(err)
		}

		hashString := app.AuthService.GenerateToken(user.Username, user.Password)
		ctx.JSON(200, modules.LoginResponse{Token: hashString})

		return nil
	}
}
