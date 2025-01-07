package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Response represents the API response structure.
type Response struct {
	AppName  string `json:"appName"`
	Language string `json:"language"`
	Version  string `json:"version"`
	PodID    string `json:"podID"`
	Env      string `json:"env"`
}

// Helper function to get an environment variable or return a default value.
func getEnv(envName, defaultValue string) string {
	if val, ok := os.LookupEnv(envName); ok {
		return val
	}
	return defaultValue
}

// Handles fetching API resource information.
func fetchAPIResourceHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		AppName:  "ML",
		Language: "golang",
		Version:  getEnv("VERSION", "0.0.0"),
		PodID:    getEnv("MY_POD_NAME", "podID_is_NOT_defined"),
		Env:      getEnv("ENV", "NOT-DEFINED"),
	}

	writeJSONResponse(w, http.StatusOK, response)
}


// Helper function to write a JSON response.
func writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Failed to encode response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// Handles health check requests.
func healthHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"status": "healthy"}
	writeJSONResponse(w, http.StatusOK, response)
}

// ML requests.
func MLHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"ML": "calculating from ML"}
	writeJSONResponse(w, http.StatusOK, response)
}

func main() {
	const port = 9992
	http.HandleFunc("/", fetchAPIResourceHandler)
	http.HandleFunc("/ml", MLHandler)
	http.HandleFunc("/health", healthHandler)
	log.Printf("Server is running on port %d...", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
