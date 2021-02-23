package main

import (
	"encoding/json"
	"fmt"
	"log"
	"persistence"
	"reflect"
	"testing"
)

func init() {
	persistence.CreateDatabase()
	fmt.Println("\nTesting customer")
}

func TestReturnAllCustomers(t *testing.T) {
	// Need to fix the ampersand in AT&T
	expected := []byte(`{"customers":[{"customerId":1,"name":"AT\u0026T"},{"customerId":2,"name":"Swisscom"}]}`)
	actual := GetCustomers()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("ReturnAllCustomers was incorrect:\nactual:   %s\nexpected: %s", actual, expected)
	}
}

func TestCreateNewCustomer(t *testing.T) {
	values := customer{Name: "Ericsson"}
	jsonData, err := json.Marshal(values)
	if err != nil {
		log.Println(err)
	}

	expected := true
	actual := CreateNewCustomer(jsonData)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("CreateNewCustomer was incorrect:\nactual:   %t\nexpected: %t", actual, expected)
	}
}

func TestGetCustomer(t *testing.T) {
	expected := []byte(`{"customerId":2,"name":"Swisscom"}`)
	actual := GetCustomer(2)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("GetCustomer was incorrect:\nactual:   %s\nexpected: %s", actual, expected)
	}
}

func TestGetCustomerFeatures(t *testing.T) {
	expected := []byte(`{"features":[{"featureId":1,"displayName":"firstFeature","technicalName":"testingFeature","expiresOn":"2022-12-31T12:00:00Z","description":"Checking if insert works","inverted":false,"active":true}]}`)
	actual := GetCustomerFeatures(2)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("GetCustomerFeatures was incorrect:\nactual:   %s\nexpected: %s", actual, expected)
	}
}
