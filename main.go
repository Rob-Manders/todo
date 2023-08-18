package main

import (
	"crypto/tls"
	"flag"
	"log"
	"net/http"
	"time"
	"todo/controllers"
	"todo/middleware"
	"todo/store"

	"github.com/gorilla/mux"
)

func main() {
	store.CreateList()

	router := mux.NewRouter()

	// Parse command line arguments:
	flagCert := flag.String("cert", "cert.pem", "path to cert")
	flagKey := flag.String("key", "key.pem", "path to key")
	flag.Parse()

	tlsConfig := createTlsConfig(flagCert, flagKey)
	
	server(router, tlsConfig)
}

func createTlsConfig(flagCert *string, flagKey *string) *tls.Config {
	// Read and parse public/private key pair. Must be PEM (Privacy-Enhanced Mail) encoded:
	certificate, error := tls.LoadX509KeyPair(*flagCert, *flagKey)
	if error != nil { log.Fatal(error) }

	// Create TLS configuration:
	return &tls.Config {
		Certificates: []tls.Certificate{certificate},
		CipherSuites: nil,
		PreferServerCipherSuites: true,
		MinVersion: tls.VersionTLS13,
		CurvePreferences: []tls.CurveID {
			tls.CurveP256,
			tls.X25519,
		},
	}
}

func server(router *mux.Router, tlsConfig *tls.Config) {
	// Redirect HTTP requests to HTTPS server:
	go func() {
		httpServer := &http.Server {
			Addr: ":3000",
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Redirect(
					w, r,
					"https://localhost:3001" + r.URL.RequestURI(),
					http.StatusMovedPermanently,
				)
			}),
		}

		log.Fatal(httpServer.ListenAndServe())
	}()

	// Initialise HTTPS server with timeouts and TLS configuration:
	httpsServer := &http.Server{
		Addr: ":3001",
		Handler: middleware.TrimPath(router),
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
