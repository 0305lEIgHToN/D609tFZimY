// 代码生成时间: 2025-10-01 17:28:39
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "log"

    "github.com/gorilla/mux"
)

// MediaAsset represents a media asset with its metadata.
type MediaAsset struct {
    ID           string `json:"id"`
    Title        string `json:"title"`
    Description  string `json:"description"`
    MediaType    string `json:"mediaType"`
    UploadDate   string `json:"uploadDate"`
}

// MediaService provides functionality to manage media assets.
type MediaService struct {
    assets map[string]MediaAsset
    nextID int
}

// NewMediaService creates a new instance of MediaService.
func NewMediaService() *MediaService {
    return &MediaService{
        assets: make(map[string]MediaAsset),
        nextID: 1,
    }
}

// AddAsset adds a new media asset to the service.
func (s *MediaService) AddAsset(title, desc, mediaType, uploadDate string) (string, error) {
    if _, exists := s.assets[fmt.Sprintf("%d", s.nextID)]; exists {
        return "", fmt.Errorf("ID %d already exists", s.nextID)
    }
    s.assets[fmt.Sprintf("%d", s.nextID)] = MediaAsset{
        ID:           fmt.Sprintf("%d", s.nextID),
        Title:        title,
        Description:  desc,
        MediaType:    mediaType,
        UploadDate:   uploadDate,
    }
    s.nextID++
    return fmt.Sprintf("%d", s.nextID-1), nil
}

// GetAsset retrieves a media asset by its ID.
func (s *MediaService) GetAsset(id string) (*MediaAsset, error) {
    asset, exists := s.assets[id]
    if !exists {
        return nil, fmt.Errorf("asset with ID %s not found", id)
    }
    return &asset, nil
}

// MediaAssetHandler handles HTTP requests for media assets.
func MediaAssetHandler(w http.ResponseWriter, r *http.Request) {
    service := NewMediaService()
    switch r.Method {
    case http.MethodGet:
        id := mux.Vars(r)["id"]
        asset, err := service.GetAsset(id)
        if err != nil {
            http.Error(w, err.Error(), http.StatusNotFound)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(asset)

    case http.MethodPost:
        var newAsset MediaAsset
        err := json.NewDecoder(r.Body).Decode(&newAsset)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        id, err := service.AddAsset(newAsset.Title, newAsset.Description, newAsset.MediaType, newAsset.UploadDate)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.WriteHeader(http.StatusCreated)
        fmt.Fprintf(w, `{"id": "%s"}`, id)

    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/assets/{id}", MediaAssetHandler).Methods("GET")
    router.HandleFunc("/assets", MediaAssetHandler).Methods("POST")
    fmt.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
