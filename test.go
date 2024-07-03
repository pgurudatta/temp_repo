/*
Sample code for vulnerable type: Improper Certificate Validation
CWE : CWE-295
Description : Improper Certificate Validation
*/
package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	url := "https://example.com"

	// Disable certificate validation (INSECURE - DO NOT USE IN PRODUCTION)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},  //source and sink
	}
	client := &http.Client{Transport: tr}

	// Send HTTP GET request to the URL
	resp, err := client.Get(url)
	if err != nil {
		log.Fatalf("Failed to send GET request: %v", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	fmt.Printf("Response Status: %s\n", resp.Status)
	fmt.Printf("Response Body: %s\n", body)
}
