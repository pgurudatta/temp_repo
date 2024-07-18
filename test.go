package main

import (
	"fmt"
	"os"
)

func main() {
	username := "admin"
	password := "s3cr3t"

	err := authenticate(username, password)
	if err != nil {
		fmt.Printf("Authentication failed: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Println("Authentication successful!")
}

func authenticate(username, password string) error {
	// Simulating authentication logic
	if username != "admin" || password != "s3cr3t" {
		return fmt.Errorf("invalid credentials for user %s", username)
	}
	return nil
}
