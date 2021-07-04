package util

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// OpenDb opens a database using the supplied file and returns the database object
func OpenDb(dbFile string) *sql.DB {
	db, err := sql.Open("sqlite3", dbFile)

	if err != nil {
		fmt.Printf("error opening db: %s\n", err.Error())
		return nil
	}

	err = db.Ping()

	if err != nil {
		fmt.Println("couldn't ping db")
		return nil
	} else {
		fmt.Println("pinged")
	}

	return db
}

// InitTables creates tables in the db if they aren't already created
func InitTables(db *sql.DB) {
	queries := []string{
		`
		CREATE TABLE IF NOT EXISTS Routines(
			id INT PRIMARY KEY,
			name TEXT,
			owner_id INT,
			collaborator_list INT,
			data TEXT
		);
		`,
		`
		CREATE TABLE IF NOT EXISTS CollaboratorLists(
			id INT PRIMARY KEY,
			routine_id INT
		);
		`,
		`
		CREATE TABLE IF NOT EXISTS Collaborators(
			user_id INT,
			collaborator_list_id INT,
			permission_level_id INT
		);
		`,
		`
		CREATE TABLE IF NOT EXISTS Users(
			id INT PRIMARY KEY,
			name TEXT
		);
		`,
		`
		CREATE TABLE IF NOT EXISTS PerimissionLevels(
			id INT PRIMARY KEY,
			description TEXT
		);
		`,
	}

	for _, query := range queries {
		initTable(db, query)
	}
}

// CloseDb closes the passed in db
func CloseDb(db *sql.DB) {
	db.Close()
}

// AddUser adds a new user to the db with the specified name
func AddUser(db *sql.DB, name string) error {
	query := `
		INSERT INTO Users VALUES(
			%s
		);
	`
	_, err := db.Exec(query, name)

	if err != nil {
		fmt.Printf("couldn't add user: %s\n", err.Error())
	}

	return err
}

// initTable runs the supplied query using the supplied db
func initTable(db *sql.DB, query string) error {
	_, err := db.Exec(query)

	if err != nil {
		fmt.Printf("couldn't create table: %s\n", err.Error())
		return err
	}

	fmt.Println("created table")
	return nil
}
