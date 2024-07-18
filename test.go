package main

import (
	"fmt"
	"os"
)

func main() {
	filename := "/etc/passwd" // Example sensitive filename
	err := openFile(filename)
	if err != nil {
		fmt.Printf("Failed to open file: %s\n", err.Error())
		os.Exit(1)
	}
}

func openFile(filename string) error {
	// Simulating an error condition where the file doesn't exist
	return fmt.Errorf("file %s not found", filename)
}
