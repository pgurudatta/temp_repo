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
		errMsg := fmt.Sprintf("Login failed for user: %s", username)
		log.Println(errMsg)
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
	}
}

func main() {
	http.HandleFunc("/login", handleLogin)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
