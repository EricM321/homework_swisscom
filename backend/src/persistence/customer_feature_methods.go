package persistence

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// GetCustomerFeatures temp
func GetCustomerFeatures(id int) []byte {
	fmt.Println("Retrieving customers features.")
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	sqlQuery := fmt.Sprintf("SELECT * FROM %s WHERE customer_id=%d", linkerTableName, id)

	stmt, err := db.Prepare(sqlQuery)
	if err != nil {
		panic(err)
	}

	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}

	var featureIDArray []int

	defer rows.Close()
	for rows.Next() {
		var customerID int
		var featureID int
		err = rows.Scan(&customerID, &featureID)
		if err != nil {
			// handle this error
			panic(err)
		}

		featureIDArray = append(featureIDArray, featureID)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	sqlQuery = fmt.Sprintf("SELECT * FROM %s WHERE feature_id=ANY(ARRAY%d)", featureTableName, featureIDArray)
	stmt, err = db.Prepare(sqlQuery)
	if err != nil {
		panic(err)
	}

	rows, err = stmt.Query()
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
