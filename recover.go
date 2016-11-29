package recover

import (
	"net/http"
)

// Version is this package's version.
const Version = "0.0.1"

// DefaultRecoverHandler is a convenient recover handler which
// simply returns "500 Internal server error".
var DefaultRecoverHandler = http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusInternalServerError)

	res.Write([]byte(http.StatusText(http.StatusInternalServerError)))
})

// Handler wraps the http.Handler with panic recovery support.
func Handler(h http.Handler, recoverHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				recoverHandler.ServeHTTP(res, req)
			}
		}()

		h.ServeHTTP(res, req)
	})
}
