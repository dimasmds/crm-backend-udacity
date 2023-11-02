package data

import (
	"github.com/carlmjohnson/truthy"
	"testing"
)

func Test_ShouldReturnCustomersWhenCallGetCustomers(t *testing.T) {
	customers := GetCustomers()

	if len(customers) != 2 {
		t.Fatalf("customers should have 2 length")
	}
}

func Test_ShouldReturnErrorWhenCallGetCustomerAndGivenByNotFoundId(t *testing.T) {
	_, er := GetCustomersById("not-found")

	if !truthy.Value(er) {
		t.Fatalf("error should be defined")
	}
}

func Test_ShouldReturnSingleCustomerWhenCallGetCustomerAndGivenByValidId(t *testing.T) {
	id := "081e798e-23ca-4db2-83bb-ddc6a1ac5727"
	customer, er := GetCustomersById("081e798e-23ca-4db2-83bb-ddc6a1ac5727")

	if customer.id != id {
		t.Fatalf("is not return correct customer")
	}

	if truthy.Value(er) {
		t.Fatalf("error is should not defined")
	}
}

func Test_ShouldAddNewCustomerWhenCallAddNewCustomer(t *testing.T) {
	newCustomer := NewCustomer{
		name:      "Test User",
		role:      "Developer",
		email:     "testing@dicoding.com",
		phone:     12345,
		contacted: false,
	}

	AddNewCustomer(newCustomer)

	// Assert
	customers := GetCustomers()

	if len(customers) != 3 {
		t.Fatalf("customer is not added properly")
	}
}

func Test_ShouldReturnErrorWhenCallUpdateCustomerWithInvalidId(t *testing.T) {
	updatedCustomer := Customer{id: "not-found"}
	er := UpdateCustomer(updatedCustomer)

	if !truthy.Value(er) {
		t.Fatalf("error should be defined")
	}
}

func Test_ShouldReturnUpdateWhenCallUpdateCustomerWithValidId(t *testing.T) {
	id := "081e798e-23ca-4db2-83bb-ddc6a1ac5727"
	customer := Customer{
		id:        id,
		name:      "Dimas Updated",
		role:      "Software Developer",
		email:     "dimas@dicoding.com",
		phone:     1234,
		contacted: false,
	}

	err := UpdateCustomer(customer)

	if truthy.Value(err) {
		t.Fatalf("error should not defined")
	}

	updatedCustomer, _ := GetCustomersById(id)

	if updatedCustomer.name != customer.name {
		t.Fatalf("is not updating properly")
	}
}

func Test_ShouldReturnErrorWhenCallDeleteCustomerWithInvalidId(t *testing.T) {
	err := DeleteCustomer("not-found")

	if !truthy.Value(err) {
		t.Fatalf("error should defined")
	}
}

func Test_ShouldReturnDeleteCustomerWhenCallDeleteCustomer(t *testing.T) {
	id := "081e798e-23ca-4db2-83bb-ddc6a1ac5727"
	err := DeleteCustomer(id)

	if truthy.Value(err) {
		t.Fatalf("error should not defined")
	}

	customers := GetCustomers()

	if len(customers) != 1 {
		t.Fatalf("customers should reduce")
	}

	customer := customers[0]

	if customer.id == id {
		t.Fatalf("removal is not valid")
	}

}
