package util

import (
	"fmt"
	"net/http"
)

func CreateRoutine(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateRoutine")
	values := r.URL.Query()

	fmt.Println(values)

	db := OpenDb("./temp.db")
	defer CloseDb(db)

	InitTables(db)
	AddUser(db, "test")

	w.Write([]byte{})
}

func LoadRoutine(w http.ResponseWriter, r *http.Request) {
	fmt.Println("LoadRoutine")

	w.Write([]byte{})
}

func SaveRoutine(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SaveRoutine")

	w.Write([]byte{})
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	db := OpenDb("./temp.db")
	user, err := GetUserByName(db, "")

	if err != nil {
	}
}
