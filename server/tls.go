package server

import (
	"crypto/tls"
	"log"
)

func CreateTlsConfig(flagCert *string, flagKey *string) *tls.Config {
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