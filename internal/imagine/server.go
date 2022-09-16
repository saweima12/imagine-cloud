package imagine

import (
	"github.com/labstack/echo"
	"github.com/saweima12/imagine/internal/pkg/webdav"
)

type Server struct {
	Echo *echo.Echo
	Dav  webdav.DAVHandler
}

func New() *Server {
	// define echo
	e := echo.New()
	// define webdav
	dav := webdav.New()

	// define handler
	s := &Server{
		Echo: e,
		Dav:  dav,
	}

	return s
}

func (s *Server) Run(port string) error {
	s.Echo.Start(port)
	return nil
}
