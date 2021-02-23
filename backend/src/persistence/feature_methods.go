package persistence

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// Feature ...
type Feature struct {
	ID            int       `json:"id"`          // optional
	DisplayName   string    `json:"displayName"` // optional
	TechnicalName string    `json:"technicalName"`
	ExpiresOn     time.Time `json:"expiresOn"`   // optional
	Description   string    `json:"description"` // optional
	Inverted      bool      `json:"inverted"`
	CustomerIds   []int     `json:"customerIds"`
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

// CreateFeature ...
func CreateFeature(value []byte) bool {
	var values Feature
	json.Unmarshal(value, &values)

	fmt.Println("Adding new feature.")
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	sqlQuery := `
	INSERT INTO ` + featureTableName + ` (` + displayName + `, ` + technicalName + `, ` + expiresOn + `, ` + description + `, ` + inverted + `, ` + active + `)
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING feature_id`

	stmt, err := db.Prepare(sqlQuery)
	if err != nil {
		log.Println(err)
		return false
	}

	var featureID int
	err = stmt.QueryRow(values.DisplayName, values.TechnicalName, values.ExpiresOn, values.Description, values.Inverted, true).Scan(&featureID)
	if err != nil {
		log.Println(err)
		return false
	}

	sqlQuery = `
	INSERT INTO ` + linkerTableName + ` (customer_id, feature_id)
	VALUES ($1, $2)`

	for _, customerID := range values.CustomerIds {
		_, err = db.Exec(sqlQuery, customerID, featureID)
		if err != nil {
			log.Panicln(err)
			return false
		}
	}
	db.Close()

	return true
}

// UpdateFeature ...
func UpdateFeature(value []byte) bool {
	fmt.Println("Updating feature.")
	var values Feature
	json.Unmarshal(value, &values)

	fmt.Println("Updating feature.")
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Panicln(err)
		return false
	}

	sqlQuery := `UPDATE ` + featureTableName + ` SET display_name = $1, technical_name = $2, expires_on = $3, description = $4, inverted = $5 WHERE feature_id = $6;`

	stmt, err := db.Prepare(sqlQuery)
	if err != nil {
		log.Panicln(err)
		return false
	}

	res, err := stmt.Exec(values.DisplayName, values.TechnicalName, values.ExpiresOn, values.Description, values.Inverted, values.ID)
	rowAffected, _ := res.RowsAffected()
	if err != nil {
		log.Panicln(err)
		return false
	} else if rowAffected == 0 {
		log.Panicln("No rows updated")
		return false
	}

	// need to remove customers from feature as well
	sqlQuery = `
	INSERT INTO ` + linkerTableName + ` (customer_id, feature_id)
	VALUES ($1, $2) ON CONFLICT DO NOTHING`

	for _, customerID := range values.CustomerIds {
		_, err = db.Exec(sqlQuery, customerID, values.ID)

		if err != nil {
			log.Panicln(err)
			return false
		}
	}
	db.Close()

	return true
}

// GetFeatures temp
func GetFeatures() []byte {
	fmt.Println("Retrieving all features.")
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
	return jsonData
}

// ToggleFeature ...
func ToggleFeature(featureID int) bool {
	fmt.Println("Toggling feature.")
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err)
		return false
	}

	sqlQuery := fmt.Sprintf("UPDATE %s SET %s= NOT %s WHERE feature_id=%d", featureTableName, inverted, inverted, featureID)

	_, err = db.Exec(sqlQuery)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

// ArchiveFeature ...
func ArchiveFeature(featureID int) bool {
	fmt.Println("Archiving feature.")
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err)
		return false
	}

	sqlQuery := fmt.Sprintf("UPDATE %s SET %s= NOT %s WHERE feature_id=%d", featureTableName, active, active, featureID)

	_, err = db.Exec(sqlQuery)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
