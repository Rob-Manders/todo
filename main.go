package main

import (
	"flag"
	"todo/server"
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

	tlsConfig := server.CreateTlsConfig(flagCert, flagKey)
	
	server.Serve(router, tlsConfig)
}
