package main

import "net/http"

func Post(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not Allowed!", http.StatusMethodNotAllowed)
		}
		next(w, r)
	}
}
