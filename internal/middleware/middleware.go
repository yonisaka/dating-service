package middleware

import (
	"net/http"
)

// Handle is contract for middleware and must implement this type for http if you need middleware http request
type Handle func(r *http.Request) int

// FilterFunc is a iterator resolver in each middleware registered
func FilterFunc(r *http.Request, mfs []Handle) int {
	for _, mf := range mfs {
		if code := mf(r); code >= 400 {
			return code
		}
	}

	return http.StatusOK
}
