package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"persistence"
	"strconv"

	"github.com/gorilla/mux"
)

func returnAllFeatures(w http.ResponseWriter, r *http.Request) {
	var feats features
	json.Unmarshal(GetFeatures(), &feats)
	json.NewEncoder(w).Encode(feats)
	fmt.Println("Endpoint Hit: returnAllFeatures")
}

func returnSingleFeature(w http.ResponseWriter, r *http.Request) {
	var feats feature
	vars := mux.Vars(r)

	// move to inside GetFeature function
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Println(err)
	} else {

		json.Unmarshal(GetFeature(id), &feats)
		json.NewEncoder(w).Encode(feats)
	}
	fmt.Println("Endpoint Hit: returnSingleFeature")
}

func createNewFeature(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "%+v", string(reqBody))
	CreateNewFeature(reqBody)
}

func updateFeature(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Println(err)
	} else {
		reqBody, _ := ioutil.ReadAll(r.Body)
		fmt.Fprintf(w, "%+v", string(reqBody))
		UpdateFreature(id, reqBody)
	}
	fmt.Println("Endpoint Hit: updateFeature")
}

func toggleFeature(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Println(err)
	} else {
		ToggleFeature(id)
	}
	fmt.Println("Endpoint Hit: toggleFeature")
}

func archiveFeature(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Println(err)
	} else {
		ArchiveFeature(id)
	}
	fmt.Println("Endpoint Hit: archiveFeature")
}

func returnAllCustomers(w http.ResponseWriter, r *http.Request) {
	var custs customers
	json.Unmarshal(GetCustomers(), &custs)
	json.NewEncoder(w).Encode(custs)
	fmt.Println("Endpoint Hit: returnAllCustomers")
}

func createNewCustomer(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "%+v", string(reqBody))
	CreateNewCustomer(reqBody)
	fmt.Println("Endpoint Hit: createNewCustomer")
}

func returnSingleCustomer(w http.ResponseWriter, r *http.Request) {
	var cust customer
	vars := mux.Vars(r)

	// move to inside GetFeature function
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Println(err)
	} else {

		json.Unmarshal(GetCustomer(id), &cust)
		json.NewEncoder(w).Encode(cust)
	}
	fmt.Println("Endpoint Hit: returnSingleCustomer")
}

func returnCustomerFeatures(w http.ResponseWriter, r *http.Request) {
	var feats features
	vars := mux.Vars(r)

	// move to inside GetFeature function
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Println(err)
	} else {

		json.Unmarshal(GetCustomerFeatures(id), &feats)
		json.NewEncoder(w).Encode(feats)
	}
	fmt.Println("Endpoint Hit: returnCustomerFeatures")
}

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/api/v1/features", returnAllFeatures)
	myRouter.HandleFunc("/api/v1/feature", createNewFeature).Methods("POST")
	myRouter.HandleFunc("/api/v1/feature/{id}", updateFeature).Methods("POST")
	myRouter.HandleFunc("/api/v1/feature/{id}", returnSingleFeature)
	myRouter.HandleFunc("/api/v1/feature/{id}/toggle", toggleFeature).Methods("POST")
	myRouter.HandleFunc("/api/v1/feature/{id}/archive", archiveFeature).Methods("POST")

	myRouter.HandleFunc("/api/v1/customers", returnAllCustomers)
	myRouter.HandleFunc("/api/v1/customer", createNewCustomer).Methods("POST")
	myRouter.HandleFunc("/api/v1/customer/{id}", returnSingleCustomer)
	myRouter.HandleFunc("/api/v1/customerfeatures/{id}", returnCustomerFeatures)

	log.Fatal(http.ListenAndServe(":10000", myRouter))

}

func main() {
	persistence.CreateDatabase()
	handleRequests()
}
