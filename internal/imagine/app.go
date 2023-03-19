package imagine

import (
	"github.com/gin-gonic/gin"
	"github.com/saweima12/imagine/internal/imagine/config"
	"github.com/saweima12/imagine/internal/service"
)

type ServerApp struct {
	UserContext *config.UserContext
	Engine      *gin.Engine
	Dav         service.WebDavService
	AuthService service.UserAuthService
}

func (app *ServerApp) Run(port string) error {
	err := app.Engine.Run(port)

	if err != nil {
		return err
	}

	return nil
}
