// 代码生成时间: 2025-11-02 16:15:32
package main

import (
    "encoding/json"
    "net/http"
    "strings"

    "github.com/gorilla/mux"
)

// Achievement represents a single achievement.
type Achievement struct {
    ID          string `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
}

// achievementStore is a simple in-memory store for achievements.
// In a real-world scenario, this would be replaced by a database.
var achievementStore = []Achievement{
    {
        ID:          "1",
        Title:       "First Login",
        Description: "You have successfully logged in for the first time.",
    },
    {
        ID:          "2",
        Title:       "Welcome Aboard",
        Description: "You have completed the onboarding process.",
    },
    // Add more achievements here as needed.
}

// getAllAchievements handles HTTP requests to retrieve all achievements.
func getAllAchievements(w http.ResponseWriter, r *http.Request) {
    // Send a JSON response with all achievements.
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(achievementStore)
}

// getAchievementByID handles HTTP requests to retrieve an achievement by its ID.
func getAchievementByID(w http.ResponseWriter, r *http.Request) {
    // Extract the achievement ID from the URL.
    vars := mux.Vars(r)
    achievementID := vars["id"]

    // Find the achievement with the given ID.
    for _, achievement := range achievementStore {
        if achievement.ID == achievementID {
            // Send a JSON response with the found achievement.
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(achievement)
            return
        }
    }

    // If no achievement is found, return a 404 Not Found status.
    http.NotFound(w, r)
}

func main() {
    // Create a new router.
    router := mux.NewRouter()

    // Define routes.
    router.HandleFunc("/achievements", getAllAchievements).Methods("GET")
    router.HandleFunc("/achievements/{id}", getAchievementByID).Methods("GET")

    // Start the server.
    // Use a port of 8080 for this example, but you can change it to any available port.
    http.ListenAndServe(":8080", router)
}
