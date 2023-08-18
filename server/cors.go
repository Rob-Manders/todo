package server

import (
	"log"
	"net/http"
)

func CorsHandler(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if (r.Method == "OPTIONS") {
			log.Print("Preflight Detected:", r.Header)
			
			w.Header().Add("Connection", "keep-alive")
			w.Header().Add("Access-Control-Allow-Origin", "http://localhost:3000")
			w.Header().Add("Access-Control-Allow-Methods", "POST, OPTIONS, GET, DELETE, PUT")
			w.Header().Add("Access-Control-Allow-Headers", "content-type")
			w.Header().Add("Access-Control-Max-Age", "86400")
		} else {
			handler.ServeHTTP(w, r)
		}
	}
}