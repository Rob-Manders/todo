package server

import (
	"net/http"
)

func Headers(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubdomains")
		w.Header().Add("Content-Security-Policy", "default-src 'self'")
		w.Header().Add("X-XSS-Protection", "1; mode=block")
		w.Header().Add("X-Frame-Options", "DENY")
		w.Header().Add("Referrer-Policy", "strict-origin-when-cross-origin")

		handler.ServeHTTP(w, r)
	}
}