// 代码生成时间: 2025-10-18 09:50:28
// supply_chain_traceability.go
package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

// Product represents a product in the supply chain.
type Product struct {
    ID          string `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
    BatchNumber string `json:"batch_number"`
   溯源信息
}

// ProductService defines the interface for product-related operations.
type ProductService interface {
    GetProduct(id string) (*Product, error)
}

// InMemoryProductService is a simple in-memory implementation of ProductService.
type InMemoryProductService struct {
    products map[string]Product
}

// NewInMemoryProductService creates a new instance of InMemoryProductService.
func NewInMemoryProductService() *InMemoryProductService {
    return &InMemoryProductService{
        products: make(map[string]Product),
    }
}

// GetProduct retrieves a product by its ID.
func (s *InMemoryProductService) GetProduct(id string) (*Product, error) {
    if product, exists := s.products[id]; exists {
        return &product, nil
    }
    return nil, fmt.Errorf("product with ID %s not found", id)
}

// InitializeProducts populates the in-memory product store with sample data.
func (s *InMemoryProductService) InitializeProducts() {
    s.products["1"] = Product{
        ID:          "1",
        Name:        "Product A",
        Description: "This is product A.",
        BatchNumber: "123456",
    }
    // Add more products as needed.
}

// ProductController handles HTTP requests related to products.
type ProductController struct {
    service ProductService
}

// NewProductController creates a new ProductController instance.
func NewProductController(service ProductService) *ProductController {
    return &ProductController{
        service: service,
    }
}

// GetProduct handles GET requests to retrieve a product by its ID.
func (c *ProductController) GetProduct(w http.ResponseWriter, r *http.Request) {
    var id string
    // Extract the product ID from the request URL.
    params := mux.Vars(r)
    id = params["id"]

    product, err := c.service.GetProduct(id)
    if err != nil {
        // Handle error by sending an appropriate response.
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    // Send the product details in JSON format.
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(product)
}

func main() {
    // Create a new router.
    router := mux.NewRouter()

    // Create an instance of the in-memory product service.
    service := NewInMemoryProductService()
    // Initialize products with sample data.
    service.InitializeProducts()

    // Create a product controller using the product service.
    controller := NewProductController(service)

    // Register the route for getting a product by its ID.
    router.HandleFunc("/products/{id}", controller.GetProduct).Methods("GET")

    // Start the HTTP server.
    fmt.Println("Starting the server on port 8080")
    if err := http.ListenAndServe(":8080", router); err != nil {
        fmt.Printf("Error starting server: %s", err)
    }
}