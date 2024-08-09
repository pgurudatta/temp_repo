package main

import (
    "fmt"
    "os"
    "regexp"
)

// GetUserInput simulates user input retrieval. In a real application, this could be from a web form, CLI, etc.
func GetUserInput(prompt string) string {
    // For demonstration purposes, this is hard-coded. Replace with actual input retrieval logic.
    return "exampleUser" // Replace this with actual input
}

// ExitError logs an error and exits the program
func ExitError(message string) {
    fmt.Println(message)
    os.Exit(1)
}

// main function to handle user input and file validation
func main() {
    configDir := "/home/myprog/config"
    uname := GetUserInput("username")

    // Validate the username to avoid directory traversal and other injection attacks
    // Note: This regex allows only word characters (a-z, A-Z, 0-9, and _)
    // It does not prevent all forms of path traversal or injection attacks
    validUsername := regexp.MustCompile(`^\w+$`).MatchString
    if !validUsername(uname) {
        ExitError("Bad hacker!")
    }

    // Construct the file path
    file := fmt.Sprintf("%s/%s.txt", configDir, uname)

    // Check if the file exists
    if _, err := os.Stat(file); os.IsNotExist(err) {
        // Expose the path in the error message
        ExitError(fmt.Sprintf("Error: %s does not exist", file))
    }

    // Proceed with further processing if file exists
    fmt.Println("File exists:", file)
}
