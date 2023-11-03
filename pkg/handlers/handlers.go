package handlers

import (
	"encoding/json"
	"github.com/dimasmds/crm-backend-udacity/pkg/data"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"os"
	"path"
)

func responseNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	err := json.NewEncoder(w).Encode(map[string]string{"message": "consumer not found"})

	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

func GetCustomersHandler(w http.ResponseWriter, _ *http.Request) {
	consumers := data.GetCustomers()

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(consumers)

	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

func GetCustomerByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	consumer, err := data.GetCustomersById(id)

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)

	if err != nil {
		responseNotFound(w)
		return
	}

	er := encoder.Encode(consumer)

	if er != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

func PostCustomerHandler(w http.ResponseWriter, r *http.Request) {
	var customer data.NewCustomer
	requestBody, _ := io.ReadAll(r.Body)
	w.Header().Set("Content-Type", "application/json")

	err := json.Unmarshal(requestBody, &customer)

	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}

	data.AddNewCustomer(customer)
	w.WriteHeader(http.StatusCreated)

	consumers := data.GetCustomers()
	err = json.NewEncoder(w).Encode(consumers)

	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
}

func UpdateCustomerByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var customer data.Customer
	w.Header().Set("Content-Type", "application/json")

	requestBody, _ := io.ReadAll(r.Body)

	err := json.Unmarshal(requestBody, &customer)

	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}

	customer.Id = id

	err = data.UpdateCustomer(customer)

	if err != nil {
		responseNotFound(w)
		return
	}

	consumers := data.GetCustomers()
	err = json.NewEncoder(w).Encode(consumers)

	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

func DeleteCustomerByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	w.Header().Set("Content-Type", "application/json")

	err := data.DeleteCustomer(id)

	if err != nil {
		responseNotFound(w)
		return
	}

	consumers := data.GetCustomers()
	err = json.NewEncoder(w).Encode(consumers)

	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

func GetStaticFileHandler() http.Handler {
	dir, _ := os.Getwd()
	p := path.Join(dir, "static")
	return http.FileServer(http.Dir(p))
}
