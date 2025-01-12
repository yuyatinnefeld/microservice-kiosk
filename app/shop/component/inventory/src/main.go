package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "strconv"
    "strings"
)

// Response represents the API response structure.
type Response struct {
    AppName  string `json:"appName"`
    Language string `json:"language"`
    Version  string `json:"version"`
    PodID    string `json:"podID"`
    Env      string `json:"env"`
}

// Item represents the structure of an item in the inventory.
type Item struct {
    Item          int     `json:"item"`
    ProductName   string  `json:"productname"`
    PurchasePrice float64 `json:"purchasePrice"`
    Price         float64 `json:"price"`
    Stock         int     `json:"stock"`
}

// Inventory holds predefined items.
var inventory = map[int]Item{
    1: {Item: 1, ProductName: "Snickers", PurchasePrice: 1.00, Price: 2.00, Stock: 100},
    2: {Item: 2, ProductName: "Hanuta Riegel Milch  Nuss - 5 Riegel", PurchasePrice: 2.00, Price: 4.00, Stock: 50},
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
        AppName:  "inventory",
        Language: "golang",
        Version:  getEnv("VERSION", "0.0.0"),
        PodID:    getEnv("MY_POD_NAME", "podID_is_NOT_defined"),
        Env:      getEnv("ENV", "NOT-DEFINED"),
    }

    writeJSONResponse(w, http.StatusOK, response)
}

// Handles creating API resources by accepting JSON input.
func createAPIResourceHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var data map[string]interface{}
    if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
        http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
        return
    }
    defer r.Body.Close()

    log.Printf("Received data: %+v\n", data)

    response := map[string]string{"message": "Data received successfully"}
    writeJSONResponse(w, http.StatusOK, response)
}

// Handles returning the details of an item based on its ID.
func getItemHandler(w http.ResponseWriter, r *http.Request) {
    // Extract item ID from the URL.
    path := strings.TrimPrefix(r.URL.Path, "/items/")
    itemID, err := strconv.Atoi(path)
    if err != nil {
        http.Error(w, "Invalid item ID", http.StatusBadRequest)
        return
    }

    // Retrieve the item from the inventory.
    item, exists := inventory[itemID]
    if !exists {
        http.Error(w, "Item not found", http.StatusNotFound)
        return
    }

    writeJSONResponse(w, http.StatusOK, item)
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
	response := map[string]string{"status": "healthy", "app": "inventory"}
    writeJSONResponse(w, http.StatusOK, response)
}

// Fetches data from an external ML API.
func fetchML(w http.ResponseWriter, r *http.Request) {
    url := "http://localhost:9992/ml"
    resp, err := http.Get(url)
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
	
	// Concatenate the string value of the "ML" key
	if mlValue, ok := jsonResponse["ML"].(string); ok {
		jsonResponse["ML"] = mlValue + " + INVENTOR - UPDATED"
	}
	
	writeJSONResponse(w, http.StatusOK, jsonResponse)
}

func main() {
    const port = 9991
    http.HandleFunc("/", fetchAPIResourceHandler)
    http.HandleFunc("/post", createAPIResourceHandler)
    http.HandleFunc("/items/", getItemHandler) // Handle all /items/<id> requests dynamically
    http.HandleFunc("/health", healthHandler)
    http.HandleFunc("/ml", fetchML)

    log.Printf("Server is running on port %d...", port)
    if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}