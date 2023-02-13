package services

import (
	"net/http"
	"regexp"

	"golang.org/x/net/webdav"
)

type WebDAVHandler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type davHandler struct {
	Native webdav.Handler
}

func NewDAV() WebDAVHandler {
	// define webdav handler from golang defualt packaage.
	native := webdav.Handler{
		FileSystem: webdav.Dir("./data"),
		LockSystem: webdav.NewMemLS(),
	}

	return &davHandler{
		Native: native,
	}
}

func (handler *davHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//  remove request's URL, prefix
	pattern := regexp.MustCompile("^/webdav")
	newPath := pattern.ReplaceAllString(req.URL.Path, "")
	req.URL.Path = newPath

	// pass to webdav handler.
	handler.Native.ServeHTTP(w, req)
}
