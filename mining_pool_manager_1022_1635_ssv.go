// 代码生成时间: 2025-10-22 16:35:54
package main

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
)

// MiningPool represents a mining pool with its properties.
type MiningPool struct {
    ID         int    `json:"id"`
    Name       string `json:"name"`
    Capacity   int    `json:"capacity"`
    CurrentMiners int `json:"currentMiners"`
}

// MiningPoolManager handles the creation and management of mining pools.
type MiningPoolManager struct {
    pools map[int]MiningPool
    nextID int
}

// NewMiningPoolManager creates a new manager for mining pools.
func NewMiningPoolManager() *MiningPoolManager {
    return &MiningPoolManager{
        pools: make(map[int]MiningPool),
        nextID: 1,
    }
}

// CreatePool adds a new mining pool to the manager.
func (m *MiningPoolManager) CreatePool(name string, capacity int) (*MiningPool, error) {
    if _, exists := m.pools[m.nextID]; exists {
        return nil, ErrPoolAlreadyExists
    }
    pool := MiningPool{
        ID:         m.nextID,
        Name:       name,
        Capacity:   capacity,
        CurrentMiners: 0,
    }
    m.pools[m.nextID] = pool
    m.nextID++
    return &pool, nil
}

// GetPool retrieves a mining pool by its ID.
func (m *MiningPoolManager) GetPool(id int) (*MiningPool, error) {
    pool, exists := m.pools[id]
    if !exists {
        return nil, ErrPoolNotFound
    }
    return &pool, nil
}

// ListPools returns a list of all mining pools.
func (m *MiningPoolManager) ListPools() []MiningPool {
    var pools []MiningPool
    for _, pool := range m.pools {
        pools = append(pools, pool)
    }
    return pools
}

// ErrPoolAlreadyExists is returned when a pool with the same ID already exists.
var ErrPoolAlreadyExists = errors.New("pool already exists")

// ErrPoolNotFound is returned when no pool with the given ID is found.
var ErrPoolNotFound = errors.New("pool not found")

// setupRoutes sets up the routes for the mining pool manager.
func setupRoutes(r *mux.Router, manager *MiningPoolManager) {
    r.HandleFunc("/pools", func(w http.ResponseWriter, r *http.Request) {
        pools, _ := manager.ListPools()
        json.NewEncoder(w).Encode(pools)
    }).Methods("GET")
    
    r.HandleFunc("/pools", func(w http.ResponseWriter, r *http.Request) {
        var pool MiningPool
        if err := json.NewDecoder(r.Body).Decode(&pool); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        _, err := manager.CreatePool(pool.Name, pool.Capacity)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.WriteHeader(http.StatusCreated)
    }).Methods("POST")
    
    r.HandleFunc("/pools/{id}", func(w http.ResponseWriter, r *http.Request) {
        var pool MiningPool
        var err error
        vars := mux.Vars(r)
        id, _ := strconv.Atoi(vars["id"])
        if pool, err = manager.GetPool(id); err != nil {
            http.Error(w, err.Error(), http.StatusNotFound)
            return
        }
        json.NewEncoder(w).Encode(pool)
    }).Methods("GET")
}

func main() {
    manager := NewMiningPoolManager()
    r := mux.NewRouter()
    setupRoutes(r, manager)
    http.ListenAndServe(":8080", r)
}
