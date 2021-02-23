package persistence

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	// Just needed for drivers of postgres in sql import
	_ "github.com/lib/pq"
)

/*
type FetureToggle struct {
	ID            int      `json:"id,omitemp"`
	DisplayName   string   `json:"displayName"`
	TechnicalName string   `json:"technicalName"`
	ExpiresOn     DateTime `json:"expiresOn"`
	Description   string   `json:"description"`
	Inverted      bool     `json:"inverted"`
	CustomerIds   []string `json:"customerIds"`
}*/

const (
	host              = "0.0.0.0"
	port              = "8080"
	user              = "postgres"
	password          = "password"
	dbName            = "feature_toggle_db"
	featureTableName  = "features"
	customerTableName = "customers"
	linkerTableName   = "customer_features"
	displayName       = "display_name"
	technicalName     = "technical_name"
	expiresOn         = "expires_on"
	description       = "description"
	active            = "active"
	inverted          = "inverted"
	customerEmail     = "email"
	customerName      = "name"
	customerRole      = "role"
)

// CreateDatabase is exported for now
func CreateDatabase() {
	fmt.Println("Creating database in postgreSQL.")
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", host, port, user, password)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	fmt.Println("Dropping database due to new create...")
	_, err = db.Exec("DROP DATABASE IF EXISTS " + dbName)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE DATABASE " + dbName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database created successfully.")
	db.Close()
	createTables()
	addInitialData()
}

func createTables() {
	fmt.Println("Adding tables to database. DB name: " + dbName)
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	// TODO remove inverted and from feature struct too
	_, err = db.Exec("CREATE TABLE " + featureTableName +
		"( feature_id serial primary key, " + displayName + " varchar(50), " + technicalName + " varchar(50) not null, " + expiresOn + " timestamp with time zone, " + description + " varchar(255), " + inverted + " boolean not null, " + active + " boolean not null)")
	if err != nil {
		log.Fatal(err)
	}

	// TODO add inverted and add to customer struct
	_, err = db.Exec("CREATE TABLE " + customerTableName +
		"( customer_id serial primary key, " + customerName + " varchar(50) not null, UNIQUE(name))")

	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE " + linkerTableName +
		"( customer_id integer not null, feature_id integer not null, primary key (customer_id, feature_id), foreign key (customer_id) references " + customerTableName + " (customer_id), foreign key (feature_id) references " + featureTableName + "(feature_id))")
	if err != nil {
		panic(err)
	}

	fmt.Println("Adding tables success.")
	db.Close()
}

func addInitialData() {
	fmt.Println("Adding data to tables.")
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	sqlQuery := `
	INSERT INTO ` + featureTableName + ` (` + displayName + `, ` + technicalName + `, ` + expiresOn + `, ` + description + `, ` + inverted + `, ` + active + `)
	VALUES ($1, $2, $3, $4, $5, $6)  RETURNING feature_id`

	stmt, err := db.Prepare(sqlQuery)
	if err != nil {
		panic(err)
	}

	var featureID int
	err = stmt.QueryRow("firstFeature", "testingFeature", time.Date(2022, 12, 31, 12, 0, 0, 0, time.UTC), "Checking if insert works", false, true).Scan(&featureID)
	if err != nil {
		panic(err)
	}

	sqlQuery = `
	INSERT INTO ` + customerTableName + ` (` + customerName + `)
	VALUES ($1)  RETURNING customer_id`

	stmt, err = db.Prepare(sqlQuery)
	if err != nil {
		panic(err)
	}

	var customerID int

	err = stmt.QueryRow("AT&T").Scan(&customerID)
	if err != nil {
		panic(err)
	}

	err = stmt.QueryRow("Swisscom").Scan(&customerID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("feature_id: %d customer_id: %d\n", featureID, customerID)

	sqlQuery = `
	INSERT INTO ` + linkerTableName + ` (customer_id, feature_id)
	VALUES ($1, $2)`

	_, err = db.Exec(sqlQuery, customerID, featureID)
	if err != nil {
		panic(err)
	}

	db.Close()
	fmt.Println("Adding data success.")
}
