package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// CreateInsecureHTTPClient creates an insecure HTTP client.
func CreateInsecureHTTPClient() *http.Client {
	// Disable certificate validation (INSECURE - DO NOT USE IN PRODUCTION)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	return client
}

// SendHTTPRequest sends an HTTP GET request using the provided client.
func SendHTTPRequest(client *http.Client, url string) ([]byte, error) {
	// Send HTTP GET request to the URL
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to send GET request: %v", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	fmt.Printf("Response Status: %s\n", resp.Status)
	fmt.Printf("Response Body: %s\n", body)

	return body, nil
}

func main() {
	url := "https://example.com"

	// Create insecure HTTP client
	client := CreateInsecureHTTPClient()

	// Send HTTP GET request using the created client
	body, err := SendHTTPRequest(client, url)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Use the response body as needed
	_ = body // Example usage
}
