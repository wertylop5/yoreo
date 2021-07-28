package util

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// InitDb creates a database instance. Should be called before using any database util function
func InitDb(dbFile string) error {
	var err error
	if db == nil {
		db, err = openDb(dbFile)
	} else {
		fmt.Println("db already initialized")
		return nil
	}

	if err != nil {
	}
	return err
}

// OpenDb opens a database using the supplied file and returns the database object
func openDb(dbFile string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbFile)

	if err != nil {
		fmt.Printf("error opening db: %s\n", err.Error())
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		fmt.Println("couldn't ping db")
		return nil, err
	} else {
		fmt.Println("pinged")
	}

	return db, nil
}

// InitTables creates tables in the db if they aren't already created
func InitTables() error {
	return initTables(db)
}

func initTables(db *sql.DB) error {
	queries := []string{
		`
		CREATE TABLE IF NOT EXISTS Routines(
			id INTEGER PRIMARY KEY NOT NULL,
			name TEXT,
			owner_id INTEGER,
			collaborator_list INTEGER,
			data TEXT
		);
		`,
		`
		CREATE TABLE IF NOT EXISTS CollaboratorLists(
			id INTEGER PRIMARY KEY NOT NULL,
			routine_id INTEGER
		);
		`,
		`
		CREATE TABLE IF NOT EXISTS Collaborators(
			user_id INTEGER,
			collaborator_list_id INTEGER,
			permission_level_id INTEGER
		);
		`,
		`
		CREATE TABLE IF NOT EXISTS Users(
			id INTEGER PRIMARY KEY NOT NULL,
			name TEXT
		);
		`,
		`
		CREATE TABLE IF NOT EXISTS PerimissionLevels(
			id INTEGER PRIMARY KEY NOT NULL,
			description TEXT
		);
		`,
	}

	var err error
	for _, query := range queries {
		err = initTable(db, query)
		if err != nil {
			return err
		}
	}

	return nil
}

// CloseDb closes the passed in db
func CloseDb(db *sql.DB) error {
	err := db.Close()
	return err
}

// AddUser adds a new user to the db with the specified name
func AddUser(name string) error {
	return addUser(db, name)
}

func addUser(db *sql.DB, name string) error {
	query := `
		INSERT INTO Users(name) VALUES(?);
	`
	_, err := db.Exec(query, name)

	if err != nil {
		fmt.Printf("couldn't add user: %s\n", err.Error())
	}

	return err
}

// Returns the first result that matches the name. If no results, returns the error
func GetUserByName(name string) (User, error) {
	return getUserByName(db, name)
}

func getUserByName(db *sql.DB, name string) (User, error) {
	rows, err := db.Query("SELECT * from Users WHERE name = ?", name)

	if err != nil {
	}

	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)

		return User{id, name}, nil
	}

	return User{}, rows.Err()
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
