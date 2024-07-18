/*
Sample code for vulnerable type: Server-Side Request Forgery (SSRF)
CWE : CWE-918
Description : Server-Side Request Forgery (SSRF)
*/
package main
import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/fetch", handleFetch)
	http.ListenAndServe(":8080", nil)
}

func handleFetch(w http.ResponseWriter, r *http.Request) {
	// Get the URL parameter from the request
	url := r.URL.Query().Get("url")   //source

	if url == "" {
		http.Error(w, "URL parameter is required", http.StatusBadRequest)
		return
	}

	// Make an HTTP GET request to the specified URL (vulnerable to SSRF)
	resp, err := http.Get(url)   //sink
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to fetch URL: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read the response body and write it to the client's response
	body := make([]byte, 1024)
	_, err = resp.Body.Read(body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to read response body: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
