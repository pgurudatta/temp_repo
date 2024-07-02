/*
Sample code for vulnerable type: Path Traversal
CWE : CWE-23
Description : Relative Path Traversal
*/

package main

import (
        "fmt"
        "net/http"
        "io/ioutil"
        "path/filepath"
)

func main() {
        // Handle requests to the "/readfile" endpoint
        http.HandleFunc("/readfile", func(w http.ResponseWriter, r *http.Request) {
                // Extract the file name from the URL query parameters
                fileName := r.URL.Query().Get("file") //source
                // Check if the file name is empty
                if fileName == "" {
                        http.Error(w, "File name not provided", http.StatusBadRequest)
                        return
                }
                // Construct the absolute path to the file
                filePath := filepath.Join("uploads", fileName)
                // Attempt to read the file
                data, err := ioutil.ReadFile(filePath)  //sink
                if err != nil {
                        // If there's an error reading the file, return an internal server error
                        http.Error(w, "Failed to read file", http.StatusInternalServerError)
                        return
                }
                // Write the file content to the response
                fmt.Fprintf(w, "%s", data)
        })

        // Start the HTTP server and listen for incoming requests on port 8080
        http.ListenAndServe(":8080", nil)
}

