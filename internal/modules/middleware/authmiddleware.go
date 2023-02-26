package middleware

import (
	"encoding/base64"
	"strings"

	"github.com/labstack/echo"
	"github.com/saweima12/imagine/internal/modules"
)

const (
	basic = "basic"
)

func CustomBasicAuth(app *modules.ImagineApp) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			auth := ctx.Request().Header.Get(echo.HeaderAuthorization)
			username, password, err := splitBasicAuthInfo(auth)

			if err != nil {
				return err
			}

			valid := app.AuthService.CheckAuthorization(username, password)
			if !valid {
				return echo.ErrUnauthorized
			}

			return next(ctx)
		}
	}
}

func splitBasicAuthInfo(auth string) (string, string, error) {

	l := len(basic)

	if len(auth) > l+1 && strings.ToLower(auth[:l]) == basic {
		b, err := base64.StdEncoding.DecodeString(auth[l+1:])
		if err != nil {
			return "", "", echo.ErrUnauthorized
		}
		cred := string(b)
		for i := 0; i < len(cred); i++ {
			if cred[i] == ':' {
				return cred[:i], cred[i+1:], nil
			}
		}
	}

	return "", "", echo.ErrUnauthorized
}
