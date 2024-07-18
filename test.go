/*
Sample code for vulnerable type: Use of a Broken or Risky Cryptographic Algorithm
CWE : CWE-327
Description : Use of a Broken or Risky Cryptographic Algorithm
*/

package main

import (
    "crypto/md5"
    "fmt"
)

func main() {
    password := "mysecretpassword"

    // Simulating storage of hashed password
    hashedPassword := hashPassword(password)

    // Simulating login attempt
    loginAttempt := "mysecretpassword"

    // Verify login
    if checkPassword(loginAttempt, hashedPassword) {
        fmt.Println("Login successful!")
    } else {
        fmt.Println("Login failed!")
    }
}

// Function to hash the password using MD5 (vulnerable)
func hashPassword(password string) string {
    hash := md5.Sum([]byte(password))  //source and sink
    return fmt.Sprintf("%x", hash)
}

// Function to check if the password matches the hashed value
func checkPassword(password, hashedPassword string) bool {
    // Recompute the hash of the provided password and compare
    return hashedPassword == hashPassword(password)
}
