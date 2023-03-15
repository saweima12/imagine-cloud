package handler

import (
	"github.com/labstack/echo"
	"github.com/saweima12/imagine/internal/imagine"
	"github.com/saweima12/imagine/internal/imagine/config"
	"github.com/saweima12/imagine/internal/imagine/middleware"
	"github.com/saweima12/imagine/internal/service"
)

func (r *Router) initApiHandler(parent *echo.Group) {
	// initialize service.
	userService := service.NewUserAuthService(*r.App.UserContext)

	// initialize handler
	handler := &apiHandler{UserAuthService: userService}

	// Add Login & Operate Endpoint.
	parent.POST("/login", handler.userlogin)

	operateGroup := parent.Group("")
	operateGroup.Use(middleware.CustomBasicAuth(userService))
	operateGroup.GET("/health", handler.checkHealth)
	operateGroup.GET("/sysInfo", handler.getSysInfo)
}

type apiHandler struct {
	UserAuthService service.UserAuthService
}

func (h *apiHandler) userlogin(ctx echo.Context) error {
	// Bind userContext to model
	user := new(config.UserContext)
	if err := ctx.Bind(user); err != nil {
		return ctx.String(400, "Bad Request.")
	}

	token, err := h.UserAuthService.VerifyUser(user.Username, user.Password)

	if err != nil {
		ctx.String(401, "Username or password Error.")
		return err
	}

	ctx.JSON(200, imagine.LoginResponse{Token: token})

	return nil

}

func (h *apiHandler) getSysInfo(ctx echo.Context) error {
	info := imagine.ReadDiskInfo("/")
	ctx.JSON(200, imagine.DashBoardResponse{
		Disk: info,
	})
	return nil
}

func (h *apiHandler) checkHealth(ctx echo.Context) error {
	ctx.String(200, "ok")
	return nil
}
