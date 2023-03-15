package imagine

import (
	"github.com/labstack/echo"
	"github.com/saweima12/imagine/internal/imagine/config"
	"github.com/saweima12/imagine/internal/service"
)

type ServerApp struct {
	UserContext *config.UserContext
	Engine      *echo.Echo
	Dav         service.WebDavService
	AuthService service.UserAuthService
}

func (app *ServerApp) Run(port string) error {
	err := app.Engine.Start(port)

	if err != nil {
		return err
	}

	return nil
}
