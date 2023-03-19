package service

import (
	"net/http"

	"golang.org/x/net/webdav"
)

type WebDavService interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type webDAVService struct {
	Native webdav.Handler
}

func NewWebDavService() WebDavService {
	// define webdav handler from golang defualt packaage.
	native := webdav.Handler{
		FileSystem: webdav.Dir("./data"),
		LockSystem: webdav.NewMemLS(),
	}

	return &webDAVService{
		Native: native,
	}
}

func (service *webDAVService) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// pass to webdav handler.
	service.Native.ServeHTTP(w, req)
}
