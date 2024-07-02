/*
Sample code for vulnerable type: Clear Text Logging
CWE : CWE-200, CWE-312
Description : Exposure of Sensitive Information to an Unauthorized Actor, Cleartext Storage of Sensitive Information
*/
package main

import (
	"fmt"
	"log"
)

type User struct {
	ID       string
	Password string
}
var userDatabase = map[string]User{
	"john_doe": {"john_doe", "secretpassword123"},
	"jane_doe": {"jane_doe", "anothersecretpassword"},
}

func main() {
	userID := "john_doe"
        password := "secretpassword123" //source
	// Authenticate user
	if authenticate(userID, password) {
		fmt.Println("Login successful!")
	} else {
		fmt.Println("Login failed. Invalid credentials.")
	}
}

func authenticate(userID, password string) bool {
	log.Printf("Attempting login with userID: %s, password: %s", userID, password) //sink
	user, ok := userDatabase[userID]
	if !ok {
		// User not found
		return false
	}

	// Check if the provided password matches the stored password
	return user.Password == password
}
