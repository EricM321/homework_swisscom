package persistence

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type customer struct {
	ID    int    `json:"customerId"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Role  string `json:"role"`
}

type customers struct {
	Customers []customer `json:"customers"`
}

type feature struct {
	ID            int       `json:"featureId"`
	DisplayName   string    `json:"displayName"`
	TechnicalName string    `json:"technicalName"`
	ExpiresOn     time.Time `json:"expiresOn"`
	Description   string    `json:"description"`
	Inverted      bool      `json:"inverted"`
	Active        bool      `json:"active"`
}

type features struct {
	Features []feature `json:"features"`
}

type relation struct {
	CustomerID int `json:"customerId"`
	FeatureID  int `json:"featureId"`
}

type relations struct {
	Relations []relation `json:"relations"`
}

// GetCustomers ...
func GetCustomers() []byte {
	fmt.Println("Retrieving customers.")
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	sqlQuery := `SELECT * FROM ` + customerTableName

	stmt, err := db.Prepare(sqlQuery)
	if err != nil {
		panic(err)
	}

	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}

	var customersArray []customer

	defer rows.Close()
	for rows.Next() {
		var id int
		var email string
		var name string
		var role string
		err = rows.Scan(&id, &email, &name, &role)
		if err != nil {
			// handle this error
			panic(err)
		}

		cust := customer{
			ID:    id,
			Email: email,
			Name:  name,
			Role:  role,
		}
		customersArray = append(customersArray, cust)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	db.Close()

	cust2 := customers{
		Customers: customersArray,
	}

	var jsonData []byte
	jsonData, err = json.Marshal(cust2)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(jsonData))
	return jsonData
}

// GetFeatures temp
func GetFeatures() []byte {
	fmt.Println("Retrieving features.")
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	sqlQuery := `SELECT * FROM ` + featureTableName

	stmt, err := db.Prepare(sqlQuery)
	if err != nil {
		panic(err)
	}

	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}

	var featuresArray []feature

	defer rows.Close()
	for rows.Next() {
		var id int
		var displayName string
		var technicalName string
		var expiresOn time.Time
		var description string
		var inverted bool
		var active bool
		err = rows.Scan(&id, &displayName, &technicalName, &expiresOn, &description, &inverted, &active)
		if err != nil {
			// handle this error
			panic(err)
		}

		feat := feature{
			ID:            id,
			DisplayName:   displayName,
			TechnicalName: technicalName,
			ExpiresOn:     expiresOn,
			Description:   description,
			Inverted:      inverted,
			Active:        active,
		}
		featuresArray = append(featuresArray, feat)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	db.Close()

	feat2 := features{
		Features: featuresArray,
	}

	var jsonData []byte
	jsonData, err = json.Marshal(feat2)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(jsonData))
	return jsonData
}

// UpdateCustomer temp
func UpdateCustomer(customerID int, role string) {

	fmt.Println("Updating customer role.")
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	sqlQuery := fmt.Sprintf("UPDATE %s SET %s='%s' WHERE customer_id=%d", customerTableName, "role", role, customerID)

	_, err = db.Exec(sqlQuery)
	if err != nil {
		panic(err)
	}

}

// UpdateFeature temp
func UpdateFeature(featureID int, invert bool) {
	fmt.Println("Updating feature toggle.")
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	sqlQuery := fmt.Sprintf("UPDATE %s SET %s='%t' WHERE feature_id=%d", featureTableName, inverted, invert, featureID)

	_, err = db.Exec(sqlQuery)
	if err != nil {
		panic(err)
	}
}

// GetCustomerFeatures temp
func GetCustomerFeatures() []byte {
	fmt.Println("Retrieving customer and feature relations.")
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	sqlQuery := `SELECT * FROM ` + linkerTableName

	stmt, err := db.Prepare(sqlQuery)
	if err != nil {
		panic(err)
	}

	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}

	var relationsArray []relation

	defer rows.Close()
	for rows.Next() {
		var customerID int
		var featureID int
		err = rows.Scan(&customerID, &featureID)
		if err != nil {
			// handle this error
			panic(err)
		}

		relat := relation{
			CustomerID: customerID,
			FeatureID:  featureID,
		}
		relationsArray = append(relationsArray, relat)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	db.Close()

	relat2 := relations{
		Relations: relationsArray,
	}

	var jsonData []byte
	jsonData, err = json.Marshal(relat2)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(jsonData))
	return jsonData
}

func archiveFeature(featureID int) {
	fmt.Println("Updating feature toggle.")
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	sqlQuery := fmt.Sprintf("UPDATE %s SET %s='%t' WHERE feature_id=%d", featureTableName, active, false, featureID)

	_, err = db.Exec(sqlQuery)
	if err != nil {
		panic(err)
	}
}
