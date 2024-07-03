/*
Sample code for vulnerable type: Insecure TLS Configuration
CWE : CWE-327
Description : Insecure TLS Configuration
*/
package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

func main() {
	// Vulnerable: Using insecure protocol and cipher suite
	config := &tls.Config{
		MinVersion: tls.VersionSSL30, // source and sink
		CipherSuites: []uint16{
			tls.TLS_RSA_WITH_RC4_128_MD5, // Insecure!
		},
	}

	// Create an HTTP server with the provided TLS configuration
	server := &http.Server{
		Addr:      ":8080",
		TLSConfig: config,
	}

	// Define a handler function for the root path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, TLS!\n")
	})

	// Start the server with TLS enabled
	err := server.ListenAndServeTLS("cert.pem", "key.pem")
	if err != nil {
		fmt.Printf("Failed to start server: %s\n", err)
		return
	}
}
