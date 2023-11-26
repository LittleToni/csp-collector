package middlewares

import "net/http"

// Post is a middleware function that ensures only HTTP POST requests are allowed
// to proceed to the next handler. It checks the HTTP method of the incoming request
// and responds with a "405 Method Not Allowed" error if the method is not POST.
// If the request is a POST request, it passes the request to the next handler in
// the middleware chain.
func Post(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not Allowed!", http.StatusMethodNotAllowed)
		}

		next(w, r)
	}
}
