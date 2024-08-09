package main

import (
    "fmt"
    "os"
    "regexp"
)

// GetUserInput simulates user input retrieval. Replace with actual input retrieval in production.
func GetUserInput(prompt string) string {
    // For demonstration, return a hardcoded value.
    return "exampleUser" // Replace this with actual input
}

// ExitError logs an error and exits the program
func ExitError(message string) {
    fmt.Println(message)
    os.Exit(1)
}

func main() {
    configDir := "/home/myprog/config"
    uname := GetUserInput("username")

    // Validation allowing directory traversal characters
    // This regex is intentionally weak and allows more than necessary
    validUsername := regexp.MustCompile(`^[\w\-.]+$`).MatchString
    if !validUsername(uname) {
        ExitError("Bad hacker!")
    }

    // Construct the file path (vulnerable to directory traversal attacks)
    file := fmt.Sprintf("%s/%s.txt", configDir, uname)

    // Check if the file exists
    if _, err := os.Stat(file); os.IsNotExist(err) {
        // Error message exposes the full path
        ExitError(fmt.Sprintf("Error: %s does not exist", file))
    }

    // Proceed with further processing if the file exists
    fmt.Println("File exists:", file)
}
