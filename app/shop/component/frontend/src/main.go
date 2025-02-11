package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

// Response represents the API response structure.
type Response struct {
	AppName  string `json:"appName"`
	Language string `json:"language"`
	Version  string `json:"version"`
	PodID    string `json:"podID"`
	Env      string `json:"env"`

}

// Item represents the structure of an item from the inventory.
type Item struct {
	Item          int     `json:"item"`
	ProductName   string  `json:"productname"`
	PurchasePrice float64 `json:"purchasePrice"`
	Price         float64 `json:"price"`
	Stock         int     `json:"stock"`
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
		AppName:  "frontend",
		Language: "golang",
		Version:  getEnv("VERSION", "0.0.0"),
		PodID:    getEnv("MY_POD_NAME", "NOT-RUNNING"),
		Env:      getEnv("ENV", "NOT-DEFINED"),
	}

	writeJSONResponse(w, http.StatusOK, response)
}

// Handles health check requests.
func healthHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"status": "healthy", "app": "frontend"}
	writeJSONResponse(w, http.StatusOK, response)
}

// Handles fetching item details from the backend with detailed logging.
func fetchItemHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("Starting fetchItemHandler")

    // Extract item index from the query parameters
    itemIndex := r.URL.Query().Get("index")
    if itemIndex == "" {
        log.Println("Missing 'index' query parameter")
        http.Error(w, "Missing 'index' query parameter", http.StatusBadRequest)
        return
    }
    log.Printf("Received 'index' query parameter: %s", itemIndex)

    // Validate item index
    if _, err := strconv.Atoi(itemIndex); err != nil {
        log.Printf("Invalid 'index' query parameter: %s", itemIndex)
        http.Error(w, "Invalid 'index' query parameter", http.StatusBadRequest)
        return
    }

    // Determine backend URL based on the environment
    env := os.Getenv("ENV")
    backendURL := determineBackendURL(env, "items", itemIndex)
    log.Printf("Determined backend URL based on ENV='%s': %s", env, backendURL)

    // Send request to backend
    log.Println("Sending request to backend")
    resp, err := http.Get(backendURL)
    if err != nil {
        log.Printf("Error calling backend: %v", err)
        http.Error(w, "Failed to connect to backend", http.StatusInternalServerError)
        return
    }
    defer func() {
        log.Println("Closing backend response body")
        resp.Body.Close()
    }()

    // Log backend response status
    log.Printf("Received response from backend with status code: %d", resp.StatusCode)

    // Read backend response
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Printf("Error reading backend response: %v", err)
        http.Error(w, "Failed to read backend response", http.StatusInternalServerError)
        return
    }
    log.Println("Successfully read backend response")

    // Forward backend response to client
    log.Println("Forwarding backend response to client")
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(resp.StatusCode)
    if _, err := w.Write(body); err != nil {
        log.Printf("Error writing response to client: %v", err)
    }
    log.Println("Response successfully sent to client")
}

// determineBackendURL constructs the backend URL based on the environment.
func determineBackendURL(env, endpoint, itemIndex string) string {
	baseURL := map[string]string{
		"K8S-DEV":   "http://cnk-backend-inventory:9991",
		"DOCKER-DEV": "http://localhost:9991",
	}[env]

	if baseURL == "" {
		baseURL = "http://127.0.0.1:9991"
	}

	if itemIndex == "" {
		return fmt.Sprintf("%s/%s", baseURL, endpoint)
	}
	return fmt.Sprintf("%s/%s/%s", baseURL, endpoint, itemIndex)
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

// Fetches data from an external ML API.
func fetchML(w http.ResponseWriter, r *http.Request) {
    env := os.Getenv("ENV")
	mlURL := determineBackendURL(env, "ml", "")

    resp, err := http.Get(mlURL)
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to fetch API response: %v", err), http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to read response body: %v", err), http.StatusInternalServerError)
        return
    }

    var jsonResponse map[string]interface{}
    if err := json.Unmarshal(body, &jsonResponse); err != nil {
        http.Error(w, fmt.Sprintf("Failed to decode JSON response: %v", err), http.StatusInternalServerError)
        return
    }

    writeJSONResponse(w, http.StatusOK, jsonResponse)
}

func main() {
	const port = 9990
	http.HandleFunc("/", fetchAPIResourceHandler)
	http.HandleFunc("/fetch-item", fetchItemHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/ml", fetchML)

	log.Printf("Server is running on port %d...", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
