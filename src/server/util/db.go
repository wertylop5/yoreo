package util

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDb() *sql.DB {
	db, err := sql.Open("sqlite3", "./temp.db")

	if err != nil {
		fmt.Printf("error opening db: %s\n", err.Error())
		return nil
	}

	//defer db.Close()

	err = db.Ping()

	if err != nil {
		fmt.Println("couldn't ping db")
		return nil
	} else {
		fmt.Println("pinged")
	}

	return db
}

func CreateTable(db *sql.DB) {
	_, err := db.Exec(
		`CREATE TABLE hi(
			id INT,
			text TEXT
		);`,
	)

	if err != nil {
		fmt.Printf("couldn't create table: %s\n", err.Error())
		return
	}
	fmt.Println("created table")
}
