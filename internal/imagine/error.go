package imagine

import (
	"fmt"
	"net/http"
)

type HTTPError struct {
	Code     int
	Msg      interface{}
	Internal error
}

var (
	ErrUnauthorized = NewHTTPError(http.StatusUnauthorized)
	ErrBadRequest   = NewHTTPError(http.StatusBadRequest)
)

func NewHTTPError(code int, msg ...interface{}) *HTTPError {
	err := &HTTPError{
		Code: code,
		Msg:  http.StatusText(code),
	}

	if len(msg) >= 1 {
		err.Msg = msg[0]
	}

	return err
}

func (err *HTTPError) Error() string {
	return fmt.Sprintf("code=%d, message=%v", err.Code, err.Msg)
}

func (err *HTTPError) Message() string {
	return fmt.Sprintf("%v", err.Msg)
}
