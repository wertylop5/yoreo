package main

import (
	"fmt"
	"net/http"
)

func main() {
	fileserv := http.FileServer(http.FileSystem(http.Dir("./dist")))

	http.Handle("/", http.StripPrefix("/", fileserv))

	fmt.Println("listening")
	http.ListenAndServe(":8000", nil)
}
