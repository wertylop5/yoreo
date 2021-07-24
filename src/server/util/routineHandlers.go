package util

import (
	"fmt"
	"net/http"
)

func CreateRoutine(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateRoutine")
	values := r.URL.Query()

	fmt.Println(values)

	AddUser("test")

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
	query := r.URL.Query()
	fmt.Println(query.Encode())
	name := query.Get("user")

	user, err := GetUserByName(name)

	if err != nil {
		w.WriteHeader(500)
		return
	}

	// w.Header().Add("Content-Type", "text/plain")
	w.Write([]byte(user.name))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
}
