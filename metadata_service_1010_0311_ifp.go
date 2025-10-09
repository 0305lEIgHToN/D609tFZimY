// 代码生成时间: 2025-10-10 03:11:24
package main

import (
    "fmt"
    "net/http"
    "log"
    "encoding/json"
    "github.com/gorilla/mux"
)

// Metadata represents a metadata entry
type Metadata struct {
    ID        string `json:"id"`
    Name      string `json:"name"`
    CreatedAt string `json:"createdAt"`
}

// metadataStore is a map to store metadata entries.
// In a real-world scenario, this would be replaced with a database.
var metadataStore = make(map[string]Metadata)

// NewMetadata creates a new metadata entry
func NewMetadata(name string) Metadata {
    return Metadata{
        ID:        generateID(),
        Name:      name,
        CreatedAt: getCurrentTime(),
    }
}

// generateID generates a unique ID for a metadata entry.
// This is a placeholder and should be replaced with a real ID generation mechanism.
func generateID() string {
    return fmt.Sprintf("%d", len(metadataStore)+1)
}

// getCurrentTime returns the current time as a string.
// This is a placeholder and should be replaced with a real time generation mechanism.
func getCurrentTime() string {
    return "2023-04-01T00:00:00Z" // Example timestamp
}

// createMetadataHandler handles the creation of metadata entries.
func createMetadataHandler(w http.ResponseWriter, r *http.Request) {
    var metadata Metadata
    err := json.NewDecoder(r.Body).Decode(&metadata)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }
    metadataStore[metadata.ID] = metadata
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(metadata)
}

// getMetadataHandler handles the retrieval of metadata entries.
func getMetadataHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    metadata, exists := metadataStore[id]
    if !exists {
        http.Error(w, "Metadata not found", http.StatusNotFound)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(metadata)
}

// deleteMetadataHandler handles the deletion of metadata entries.
func deleteMetadataHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    if _, exists := metadataStore[id]; !exists {
        http.Error(w, "Metadata not found", http.StatusNotFound)
        return
    }
    delete(metadataStore, id)
    w.WriteHeader(http.StatusOK)
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/metadata", createMetadataHandler).Methods("POST")
    router.HandleFunc("/metadata/{id}", getMetadataHandler).Methods("GET")
    router.HandleFunc("/metadata/{id}", deleteMetadataHandler).Methods("DELETE")

    log.Println("Starting metadata service on port 8080")
    err := http.ListenAndServe(":8080", router)
    if err != nil {
        log.Fatal("Error starting server: ", err)
    }
}