/*
Sample code for vulnerable type: Inadequate Encryption Strength
CWE : CWE-326
Description : Inadequate Encryption Strength
*/
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

func GenerateRSAKeyPair(bits int) (*rsa.PrivateKey, error) {
	pvk, err := rsa.GenerateKey(rand.Reader, bits)  //Sink
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
	pvk, err := GenerateRSAKeyPair(1024)  //Source
	if err != nil {
		fmt.Println(err)
		return
	}
	PrintPrivateKey(pvk)
}
