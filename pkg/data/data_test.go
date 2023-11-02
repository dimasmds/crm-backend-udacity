package data

import (
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

	if er == nil {
		t.Fatalf("error should be defined")
	}
}

func Test_ShouldReturnSingleCustomerWhenCallGetCustomerAndGivenByValidId(t *testing.T) {
	id := "081e798e-23ca-4db2-83bb-ddc6a1ac5727"
	customer, er := GetCustomersById("081e798e-23ca-4db2-83bb-ddc6a1ac5727")

	if customer.Id != id {
		t.Fatalf("is not return correct customer")
	}

	if er != nil {
		t.Fatalf("error is should not defined")
	}
}

func Test_ShouldAddNewCustomerWhenCallAddNewCustomer(t *testing.T) {
	newCustomer := NewCustomer{
		Name:      "Test User",
		Role:      "Developer",
		Email:     "testing@dicoding.com",
		Phone:     12345,
		Contacted: false,
	}

	AddNewCustomer(newCustomer)

	// Assert
	customers := GetCustomers()

	if len(customers) != 3 {
		t.Fatalf("customer is not added properly")
	}
}

func Test_ShouldReturnErrorWhenCallUpdateCustomerWithInvalidId(t *testing.T) {
	updatedCustomer := Customer{Id: "not-found"}
	er := UpdateCustomer(updatedCustomer)

	if er == nil {
		t.Fatalf("error should be defined")
	}
}

func Test_ShouldReturnUpdateWhenCallUpdateCustomerWithValidId(t *testing.T) {
	id := "081e798e-23ca-4db2-83bb-ddc6a1ac5727"
	customer := Customer{
		Id:        id,
		Name:      "Dimas Updated",
		Role:      "Software Developer",
		Email:     "dimas@dicoding.com",
		Phone:     1234,
		Contacted: false,
	}

	err := UpdateCustomer(customer)

	if err != nil {
		t.Fatalf("error should not defined")
	}

	updatedCustomer, _ := GetCustomersById(id)

	if updatedCustomer.Name != customer.Name {
		t.Fatalf("is not updating properly")
	}
}

func Test_ShouldReturnErrorWhenCallDeleteCustomerWithInvalidId(t *testing.T) {
	err := DeleteCustomer("not-found")

	if err == nil {
		t.Fatalf("error should defined")
	}
}

func Test_ShouldReturnDeleteCustomerWhenCallDeleteCustomer(t *testing.T) {
	id := "081e798e-23ca-4db2-83bb-ddc6a1ac5727"
	err := DeleteCustomer(id)

	if err != nil {
		t.Fatalf("error should not defined")
	}

	customers := GetCustomers()

	customer := customers[0]

	if customer.Id == id {
		t.Fatalf("removal is not valid")
	}

}
