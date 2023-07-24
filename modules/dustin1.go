package main

import (
	"encoding/json"
	"net/http"
)

type MyData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Simulated data for the response
	data := MyData{
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}

	// Convert data to JSON
	responseJSON, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.Write(responseJSON)
}

func main() {
	http.HandleFunc("/api/data", handleRequest)
	http.ListenAndServe(":8080", nil)
}
