package main

import (
	"persistence"
)

type customer struct {
	ID   int    `json:"customerId"`
	Name string `json:"name"`
}

type customers struct {
	Customers []customer `json:"customers"`
}

// TODO Centralise structs
type relation struct {
	CusID  int `json:"customerId"`
	FeatID int `json:"featureId"`
}

type relations struct {
	Relations []relation `json:"relations"`
}

// TODO Refactor to rest calls

// GetCustomers will give a json object with all customers in the database
func GetCustomers() []byte {
	return persistence.GetCustomers()
}

// CreateNewCustomer returns if creation was a success
func CreateNewCustomer(values []byte) bool {
	return persistence.CreateCustomer(values)
}

// GetCustomer returns the row of data for individual customer
func GetCustomer(id int) []byte {
	return persistence.GetCustomer(id)
}

// GetCustomerFeatures returns all features that are connected to give customer
func GetCustomerFeatures(id int) []byte {
	return persistence.GetCustomerFeatures(id)
}
