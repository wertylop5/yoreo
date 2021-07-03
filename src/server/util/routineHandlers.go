package util

import (
	"fmt"
	"net/http"
)

func CreateRoutine(w http.ResponseWriter, r *http.Request) {
	fmt.Println("createRoutine")

	w.Write([]byte{})
}
