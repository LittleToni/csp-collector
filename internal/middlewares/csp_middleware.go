package middlewares

import "net/http"

// CSPContentType is a middleware function for HTTP handlers that ensures the request
// has the correct Content-Type header for CSP (Content Security Policy) reports.
// It checks if the incoming HTTP request's Content-Type header is set to
// "application/csp-report". If not, it responds with a "415 Unsupported Media Type"
// error. If the Content-Type is correct, it calls the next handler in the chain.
func CSPContentType(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/csp-report" {
			http.Error(w, "Invalid Content-Type!", http.StatusUnsupportedMediaType)
		}

		next(w, r)
	}
}
