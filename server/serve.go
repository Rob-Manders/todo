package server

import (
	"crypto/tls"
	"log"
	"net/http"
	"time"
	"todo/controllers"
	"todo/middleware"

	"github.com/gorilla/mux"
)

func Serve(router *mux.Router, tlsConfig *tls.Config) {
	// Initialise HTTPS server with timeouts and TLS configuration:
	httpsServer := &http.Server{
		Addr: ":3000",
		Handler: Headers(CorsHandler(middleware.TrimPath(router))),
		TLSConfig: tlsConfig,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		
		// Other timeouts are also possible:
		// IdleTimeout: 120 * time.Second,
		// ReadHeaderTimeout: 5 * time.Second,
	}

	// Static Files
	fileServer := http.FileServer(http.Dir("static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))

	// Frontend Routes
	router.HandleFunc("/", controllers.Home)
	router.HandleFunc("/about", controllers.About)

	// API
	router.HandleFunc("/api/create", controllers.CreateItem)
	router.HandleFunc("/api/delete", controllers.DeleteItem)

	log.Println("Starting server on port 3000.")
	log.Fatal(httpsServer.ListenAndServeTLS("", ""))
}

// Redirect HTTP requests to HTTPS server:
// Doesn't work...

// go func() {
// 	httpServer := &http.Server {
// 		Addr: ":3000",
// 		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			http.Redirect(
// 				w, r,
// 				"https://localhost:3001" + r.URL.RequestURI(),
// 				http.StatusMovedPermanently,
// 			)
// 		}),
// 	}

// 	log.Fatal(httpServer.ListenAndServe())
// }()
