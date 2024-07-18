package main

import (
	"fmt"
	"os"
)

func main() {
	username := "admin"
	password := "s3cr3t"

	err := login(username, password)
	if err != nil {
		fmt.Printf("Login failed: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Println("Login successful!")
}

func login(username, password string) error {
	// Simulating authentication logic
	if username != "admin" || password != "s3cr3t" {
		return fmt.Errorf("invalid credentials for user %s", username)
	}
	return nil
}
