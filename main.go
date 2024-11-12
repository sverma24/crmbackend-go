package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type Customer struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Contacted bool   `json:"contacted"`
}

var (
	customers = []Customer{
		{Id: 1, Name: "John Doe", Role: "Admin", Email: "john@example.com", Phone: "1234567890", Contacted: false},
		{Id: 2, Name: "Jane Smith", Role: "User", Email: "jane@example.com", Phone: "0987654321", Contacted: true},
		{Id: 3, Name: "Alice Brown", Role: "Manager", Email: "alice@example.com", Phone: "1112223333", Contacted: false},
	}
	mu     sync.Mutex
	nextID = 4 // Start from 4 since you already have IDs 1, 2, and 3
)

// Helper function to find a customer by ID
func findCustomerById(id int) (*Customer, int) {
	for index, c := range customers {
		if c.Id == id {
			return &c, index
		}
	}
	return nil, -1
}

// Handler to get all customers
func getCustomers(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customers)
}

// Handler to get a specific customer by ID
func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := strings.TrimPrefix(r.URL.Path, "/customers/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid customer ID", http.StatusBadRequest)
		return
	}

	customer, _ := findCustomerById(id)
	if customer != nil {
		json.NewEncoder(w).Encode(customer)
	} else {
		http.Error(w, fmt.Sprintf("Customer with ID %d not found", id), http.StatusNotFound)
	}
}

// Handler to add a new customer
func addCustomer(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var newCustomer Customer
	if err := json.Unmarshal(body, &newCustomer); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Lock to ensure thread-safe ID assignment and slice modification
	mu.Lock()
	defer mu.Unlock()

	// Assign a unique ID automatically
	newCustomer.Id = nextID
	nextID++

	// Add the new customer to the list
	customers = append(customers, newCustomer)

	// Respond with the newly added customer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(customers)
}

// Handler to update an existing customer
func updateCustomer(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/customers/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid customer ID", http.StatusBadRequest)
		return
	}

	// Use ioutil.ReadAll to read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// Unmarshal the JSON data into a Customer struct
	var updatedCustomer Customer
	if err := json.Unmarshal(body, &updatedCustomer); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Find the existing customer by ID
	customer, index := findCustomerById(id)
	if customer != nil {
		// Update the customer fields
		updatedCustomer.Id = customer.Id // Retain the original ID
		customers[index] = updatedCustomer

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updatedCustomer)
	} else {
		http.Error(w, fmt.Sprintf("Customer with ID %d not found", id), http.StatusNotFound)
	}
}

// Handler to delete a customer by ID
func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/customers/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid customer ID", http.StatusBadRequest)
		return
	}

	_, index := findCustomerById(id)
	if index != -1 {
		customers = append(customers[:index], customers[index+1:]...)
		w.WriteHeader(http.StatusNoContent)
	} else {
		http.Error(w, fmt.Sprintf("Customer with ID %d not found", id), http.StatusNotFound)
	}
}

func main() {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./static"))

	mux.Handle("/", fileServer)

	mux.HandleFunc("/customers", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getCustomers(w)
		case "POST":
			addCustomer(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	// Using dynamic path for customer ID
	mux.HandleFunc("/customers/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getCustomer(w, r)
		case "PUT":
			updateCustomer(w, r)
		case "DELETE":
			deleteCustomer(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server is starting on port 3000...")
	http.ListenAndServe(":3000", mux)
}
