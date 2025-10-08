// 代码生成时间: 2025-10-09 02:42:23
package main

import (
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/gorilla/schema"
    "log"
    "strings"
)

// InteractionService represents the service layer for teacher-student interaction
type InteractionService struct {
    decoder *schema.Decoder
}

// NewInteractionService creates a new interaction service
func NewInteractionService() *InteractionService {
    decoder := schema.NewDecoder()
    decoder.IgnoreUnknownKeys(true)
    return &InteractionService{decoder: decoder}
}

// Interaction represents the data structure for teacher-student interactions
type Interaction struct {
    Message string `json:"message"`
    Sender  string `json:"sender"`
}

// InteractionHandler handles the teacher-student interaction request
func (s *InteractionService) InteractionHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    sender := vars["sender"]
    message := vars["message"]

    if sender == "" || message == "" {
        http.Error(w, "missing sender or message", http.StatusBadRequest)
        return
    }

    // Process the interaction (e.g., store it, send a notification, etc.)
    // For simplicity, we're just logging it here
    log.Printf("Received interaction from %s: %s", sender, message)

    // Respond with success
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

// main function to start the server
func main() {
    r := mux.NewRouter()
    service := NewInteractionService()

    // Define the route for teacher-student interaction
    r.HandleFunc("/interaction/{sender}/{message}", service.InteractionHandler).Methods("POST")

    log.Println("Server is running on port 8080")
    // Start the server
    err := http.ListenAndServe(":8080", r)
    if err != nil {
        log.Fatal(err)
    }
}
