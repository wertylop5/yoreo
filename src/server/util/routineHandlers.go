package util

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateRoutine(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(404)
		w.Write([]byte{})
	}

	fmt.Println("CreateRoutine")
	type Body = struct {
		name string
	}
	body := Body{}
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Couldn't parse request body as json"))
	}

	fmt.Println(body.name)

	AddUser(body.name)

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
	if r.Method != "POST" {
		w.WriteHeader(404)
		w.Write([]byte{})
	}
	fmt.Println("CreateUser")

	type Body = struct {
		Name string `json:"name"`
	}
	var body Body
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Couldn't parse request body as json"))
	}

	fmt.Println(body.Name)

	AddUser(body.Name)

	w.Write([]byte{})
}
