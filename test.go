package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleLogin(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Authenticate the user
	if username == "admin" && password == "secretpassword" {
		// Successful login
		fmt.Fprintf(w, "Welcome, admin!")
	} else {
		// Failed login with detailed sensitive information
		errMsg := fmt.Sprintf("Login failed for user: %s with password: %s", username, password)
		log.Println("Detailed Error Message:", errMsg) // Log sensitive information
		http.Error(w, fmt.Sprintf("Login failed for user: %s with password: %s", username, password), http.StatusUnauthorized) // Expose sensitive information in response
	}
}

func main() {
	http.HandleFunc("/login", handleLogin)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
