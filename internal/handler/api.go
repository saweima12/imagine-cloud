package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/saweima12/imagine/internal/imagine"
	"github.com/saweima12/imagine/internal/imagine/config"
	"github.com/saweima12/imagine/internal/service"
)

func (r *Router) initApiHandler(parent *gin.RouterGroup) {
	// initialize service.
	userService := service.NewUserAuthService(*r.App.UserContext)

	// initialize handler
	handler := &apiHandler{UserAuthService: userService}

	// Add Login & Operate Endpoint.
	parent.POST("/login", handler.userlogin)
	parent.GET("/sysInfo", handler.getSysInfo)
	parent.GET("/health", handler.checkHealth)
}

type apiHandler struct {
	UserAuthService service.UserAuthService
}

func (h *apiHandler) userlogin(ctx *gin.Context) {
	// Bind userContext to model
	user := new(config.UserContext)
	if err := ctx.Bind(user); err != nil {
		ctx.String(400, "Bad Request.")
		return
	}

	token, err := h.UserAuthService.VerifyUser(user.Username, user.Password)

	if err != nil {
		ctx.String(401, "Username or password Error.")
		return
	}

	ctx.JSON(200, imagine.LoginResponse{Token: token})
}

func (h *apiHandler) getSysInfo(ctx *gin.Context) {
	info := imagine.ReadDiskInfo("/")
	ctx.JSON(200, imagine.DashBoardResponse{
		Disk: info,
	})
}

func (h *apiHandler) checkHealth(ctx *gin.Context) {
	ctx.String(200, "ok")
}
