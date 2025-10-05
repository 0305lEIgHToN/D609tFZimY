// 代码生成时间: 2025-10-06 03:14:26
package main

import (
    "net/http"
    "encoding/json"
    "log"
    "github.com/gorilla/mux"
)

// ClinicalTrial represents a clinical trial
type ClinicalTrial struct {
    ID           string `json:"id"`
    Name         string `json:"name"`
    Description  string `json:"description"`
    Status       string `json:"status"`
}

// TrialService handles the logic for clinical trials
type TrialService struct {
    trials map[string]ClinicalTrial
}

// NewTrialService creates a new instance of TrialService
func NewTrialService() *TrialService {
    return &TrialService{
        trials: make(map[string]ClinicalTrial),
    }
}

// AddTrial adds a new clinical trial to the service
func (s *TrialService) AddTrial(trial ClinicalTrial) string {
    s.trials[trial.ID] = trial
    return trial.ID
}

// GetTrial retrieves a clinical trial by ID
func (s *TrialService) GetTrial(id string) (ClinicalTrial, error) {
    trial, exists := s.trials[id]
    if !exists {
        return ClinicalTrial{}, ErrTrialNotFound
    }
    return trial, nil
}

// ErrTrialNotFound is returned when a trial is not found
var ErrTrialNotFound = errors.New("clinical trial not found")

// API handles HTTP requests and responses
type API struct {
    service *TrialService
}

// NewAPI creates a new instance of API
func NewAPI(service *TrialService) *API {
    return &API{service: service}
}

// Routes defines the routes for the API
func (api *API) Routes() *mux.Router {
    router := mux.NewRouter()
    router.HandleFunc("/trials", api.addTrial).Methods("POST")
    router.HandleFunc("/trials/{id}", api.getTrial).Methods("GET")
    return router
}

// addTrial handles POST requests to create a new clinical trial
func (api *API) addTrial(w http.ResponseWriter, r *http.Request) {
    var trial ClinicalTrial
    if err := json.NewDecoder(r.Body).Decode(&trial); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    trial.ID = api.service.AddTrial(trial)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(trial)
}

// getTrial handles GET requests to retrieve a clinical trial by ID
func (api *API) getTrial(w http.ResponseWriter, r *http.Request) {
    var id = mux.Vars(r)["id"]
    trial, err := api.service.GetTrial(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(trial)
}

func main() {
    // Create a new service
    service := NewTrialService()
    // Create a new API with the service
    api := NewAPI(service)
    // Define the routes
    router := api.Routes()
    // Start the server
    log.Fatal(http.ListenAndServe(":8080", router))
}
