package persistence

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
)

/*
type user struct {
	ID    int    `json:"customerId"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Role  string `json:"role"`
}
*/

type customer struct {
	ID   int    `json:"customerId"`
	Name string `json:"name"`
}

type customers struct {
	Customers []customer `json:"customers"`
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
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			// handle this error
			panic(err)
		}

		cust := customer{
			ID:   id,
			Name: name,
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

// GetCustomer ...
func GetCustomer(id int) []byte {
	fmt.Println("Retrieving customer.")
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	sqlQuery := fmt.Sprintf("SELECT * FROM %s WHERE customer_id=%d", customerTableName, id)

	stmt, err := db.Prepare(sqlQuery)
	if err != nil {
		panic(err)
	}

	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}

	var cust customer

	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			// handle this error
			panic(err)
		}

		cust = customer{
			ID:   id,
			Name: name,
		}
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	db.Close()

	var jsonData []byte
	jsonData, err = json.Marshal(cust)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(jsonData))
	return jsonData
}

// CreateCustomer ...
func CreateCustomer(name string) bool {
	fmt.Println("Adding new customer.")
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err)
		return false
	}

	sqlQuery := `INSERT INTO ` + customerTableName + ` (` + customerName + `) VALUES ($1)`
	stmt, err := db.Prepare(sqlQuery)
	if err != nil {
		log.Println(err)
		return false
	}

	_, err = stmt.Exec(name)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

// UpdateCustomer would be update user in the future
/*func updateCustomer(customerID int, role string) {

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

}*/
