package middleware

import (
	"encoding/base64"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/saweima12/imagine/internal/imagine"
	"github.com/saweima12/imagine/internal/service"
)

const (
	HeaderAuthorization = "Authorization"
	Basic               = "basic"
)

func CustomBasicAuth(userAuthService service.UserAuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.Request.Header.Get(HeaderAuthorization)
		username, password, err := splitBasicAuthInfo(auth)

		if err != nil {
			ctx.String(err.Code, err.Message())
			return
		}

		valid := userAuthService.CheckAuthorization(username, password)
		if !valid {

			ctx.String(imagine.ErrUnauthorized.Code, imagine.ErrBadRequest.Message())
			return
		}

	}
}

func splitBasicAuthInfo(auth string) (string, string, *imagine.HTTPError) {

	l := len(Basic)

	if len(auth) > l+1 && strings.ToLower(auth[:l]) == Basic {
		b, err := base64.StdEncoding.DecodeString(auth[l+1:])
		if err != nil {
			return "", "", imagine.ErrBadRequest
		}
		cred := string(b)
		for i := 0; i < len(cred); i++ {
			if cred[i] == ':' {
				return cred[:i], cred[i+1:], nil
			}
		}
	}

	return "", "", imagine.ErrUnauthorized
}
