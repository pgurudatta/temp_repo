/*
Sample code for vulnerable type: Use of Insufficiently Random Values
CWE : CWE-330
Description : Use of Insufficiently Random Values
*/
package main

import (
	"math/rand"
)

var charset = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// Function to generate a random password
func generatePassword() string {
	s := make([]rune, 20)
	for i := range s {
		s[i] = getRandomChar()
	}
	return string(s)
}

// Function to get a random character from the charset
func getRandomChar() rune {
	return charset[rand.Intn(len(charset))] // Source
}

// Function to print the generated password
func printPassword(password string) {
	println(password) // Sink
}

func main() {
	// Generate a random password
	password := generatePassword()

	// Print the generated password
	printPassword(password)
}
