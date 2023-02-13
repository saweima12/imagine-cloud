package modules

import (
	"github.com/labstack/echo"
	"github.com/saweima12/imagine/internal/services"
)

type ImagineApp struct {
	Engine      *echo.Echo
	Dav         services.WebDAVHandler
	AuthService services.UserAuthService
}

func (app *ImagineApp) Run(port string) error {
	err := app.Engine.Start(port)

	if err != nil {
		return err
	}

	return nil
}
