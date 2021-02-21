package main

import "persistence"

func returnAllCustomers() []byte {
	// Refactor to point to rest endpoint on persistence layer
	return persistence.GetCustomers()
}

func createNewCustomer(email string, name string, role string) {

}
