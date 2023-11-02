package data

import (
	"errors"
	"github.com/google/uuid"
)

type Customer struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     int    `json:"phone"`
	Contacted bool   `json:"contacted"`
}

type NewCustomer struct {
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     int    `json:"phone"`
	Contacted bool   `json:"contacted"`
}

var customers = []Customer{
	{
		Id:        "081e798e-23ca-4db2-83bb-ddc6a1ac5727",
		Name:      "Dimas Saputra",
		Role:      "Software Engineer",
		Email:     "dimas@dicoding.com",
		Phone:     621324,
		Contacted: false,
	},
	{
		Id:        "b19e6af6-725b-4574-8aad-ab79490466aa",
		Name:      "Nina Pratiwi",
		Role:      "Artist",
		Email:     "nina@dicoding.com",
		Phone:     621321,
		Contacted: true,
	},
}

func GetCustomers() []Customer {
	return customers
}

func GetCustomersById(id string) (*Customer, error) {
	var customer Customer

	for _, c := range customers {
		if c.Id == id {
			customer = c
			return &customer, nil
		}
	}

	return nil, errors.New("customer is not found")
}

func AddNewCustomer(newCustomer NewCustomer) {
	id := uuid.New().String()
	customer := Customer{
		Id:        id,
		Name:      newCustomer.Name,
		Role:      newCustomer.Role,
		Email:     newCustomer.Email,
		Phone:     newCustomer.Phone,
		Contacted: newCustomer.Contacted,
	}

	customers = append(customers, customer)
}

func UpdateCustomer(updatedCustomer Customer) error {
	for i, c := range customers {
		if c.Id == updatedCustomer.Id {
			customers[i] = updatedCustomer
			return nil
		}
	}

	return errors.New("customer is not found")
}

func DeleteCustomer(id string) error {
	for i, c := range customers {
		if c.Id == id {
			customers = append(customers[:i], customers[i+1:]...)
			return nil
		}
	}

	return errors.New("customer is not found")
}
