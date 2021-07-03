package util

import (
	"fmt"
	"net/http"
)

func CreateRoutine(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateRoutine")

	db := OpenDb()
	CreateTable(db)

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
