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
		// Failed login
		errMsg := fmt.Sprintf("Login failed for user: %s with password: %s", username, password)
		log.Println(errMsg) // Log the detailed error message
		http.Error(w, fmt.Sprintf("Login failed for user: %s", username), http.StatusUnauthorized)
	}
}

func main() {
	http.HandleFunc("/login", handleLogin)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
