package persistence

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

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

// GetFeature ...
func GetFeature(id int) []byte {
	fmt.Println("Retrieving feature.")
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	sqlQuery := fmt.Sprintf("SELECT * FROM %s WHERE feature_id=%d", featureTableName, id)

	stmt, err := db.Prepare(sqlQuery)
	if err != nil {
		panic(err)
	}

	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}

	var feat feature

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

		feat = feature{
			ID:            id,
			DisplayName:   displayName,
			TechnicalName: technicalName,
			ExpiresOn:     expiresOn,
			Description:   description,
			Inverted:      inverted,
			Active:        active,
		}
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	db.Close()

	var jsonData []byte
	jsonData, err = json.Marshal(feat)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(jsonData))
	return jsonData
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
