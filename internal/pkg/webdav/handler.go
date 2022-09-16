package webdav

import (
	"net/http"

	"golang.org/x/net/webdav"
)

type DAVHandler interface {
	ServeHTTP(w *http.ResponseWriter, r *http.Request)
}

type handler struct {
	Native webdav.Handler
}

func New() DAVHandler {
	// define webdav handler from golang defualt packaage.
	h := webdav.Handler{
		FileSystem: webdav.Dir("./data"),
		LockSystem: webdav.NewMemLS(),
	}

	return &handler{
		Native: h,
	}
}

func (h *handler) ServeHTTP(w *http.ResponseWriter, r *http.Request) {
	h.Native.ServeHTTP(*w, r)
}
