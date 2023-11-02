package data

import (
	"errors"
	"github.com/google/uuid"
)

type Customer struct {
	id        string
	name      string
	role      string
	email     string
	phone     int
	contacted bool
}

type NewCustomer struct {
	name      string
	role      string
	email     string
	phone     int
	contacted bool
}

var customers = []Customer{
	{
		id:        "081e798e-23ca-4db2-83bb-ddc6a1ac5727",
		name:      "Dimas Saputra",
		role:      "Software Engineer",
		email:     "dimas@dicoding.com",
		phone:     621324,
		contacted: false,
	},
	{
		id:        "b19e6af6-725b-4574-8aad-ab79490466aa",
		name:      "Nina Pratiwi",
		role:      "Artist",
		email:     "nina@dicoding.com",
		phone:     621321,
		contacted: true,
	},
}

func GetCustomers() []Customer {
	return customers
}

func GetCustomersById(id string) (*Customer, error) {
	var customer Customer

	for _, c := range customers {
		if c.id == id {
			customer = c
			return &customer, nil
		}
	}

	return nil, errors.New("customer is not found")
}

func AddNewCustomer(newCustomer NewCustomer) {
	id := uuid.New().String()
	customer := Customer{
		id:        id,
		name:      newCustomer.name,
		role:      newCustomer.role,
		email:     newCustomer.email,
		phone:     newCustomer.phone,
		contacted: newCustomer.contacted,
	}

	customers = append(customers, customer)
}

func UpdateCustomer(updatedCustomer Customer) error {
	for i, c := range customers {
		if c.id == updatedCustomer.id {
			customers[i] = updatedCustomer
			return nil
		}
	}

	return errors.New("customer is not found")
}

func DeleteCustomer(id string) error {
	for i, c := range customers {
		if c.id == id {
			customers = append(customers[:i], customers[i+1])
			return nil
		}
	}

	return errors.New("customer is not found")
}
