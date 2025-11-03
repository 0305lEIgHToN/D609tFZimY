// 代码生成时间: 2025-11-03 10:27:53
package main

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
)

// Employee represents the data structure for an employee
type Employee struct {
    ID       int    "json:"id"
    Name     string "json:"name"
    Position string "json:"position"
}

// EmployeeService handles business logic for employee data
type EmployeeService struct {
    // data store for employee could be a database or in-memory
    employees map[int]Employee
}

// NewEmployeeService initializes a new EmployeeService
func NewEmployeeService() *EmployeeService {
    return &EmployeeService{
        employees: make(map[int]Employee),
    }
}

// AddEmployee adds a new employee to the service
func (es *EmployeeService) AddEmployee(w http.ResponseWriter, r *http.Request) {
    var employee Employee
    if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    es.employees[employee.ID] = employee
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(employee)
}

// GetEmployee retrieves an employee by ID
func (es *EmployeeService) GetEmployee(w http.ResponseWriter, r *http.Request) {
    var employee Employee
    var err error
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    employee, ok := es.employees[id]
    if !ok {
        http.Error(w, "Employee not found", http.StatusNotFound)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(employee)
}

func main() {
    r := mux.NewRouter()
    es := NewEmployeeService()

    // Define routes
    r.HandleFunc("/employees", es.AddEmployee).Methods("POST")
    r.HandleFunc("/employees/{id}", es.GetEmployee).Methods("GET")

    // Start the server
    http.ListenAndServe(":8080", r)
}
