package main

import (
    "fmt"
    "os"
)

func main() {
    // Simulating a function that attempts to open a file
    filename := "/etc/passwd" // Example sensitive file
    err := openFile(filename)
    if err != nil {
        fmt.Printf("Failed to open file: %s\n", err.Error())
        os.Exit(1)
    }
}

// Function to open a file (simulated)
func openFile(filename string) error {
    // Simulating an error condition where the file doesn't exist
    return fmt.Errorf("file %s not found", filename)
}
