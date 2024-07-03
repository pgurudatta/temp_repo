package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

// Source: GenerateRSAKeyPair generates an RSA private key pair.
func GenerateRSAKeyPair(bits int) (*rsa.PrivateKey, error) {
	pvk, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, fmt.Errorf("failed to generate RSA key pair: %v", err)
	}
	return pvk, nil
}

// Sink: PrintPrivateKey prints the private key to stdout.
func PrintPrivateKey(pvk *rsa.PrivateKey) {
	fmt.Println(pvk)
}

func main() {
	// Generate RSA private key pair (source)
	pvk, err := GenerateRSAKeyPair(1024)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print or use the private key (sink)
	PrintPrivateKey(pvk)
}
