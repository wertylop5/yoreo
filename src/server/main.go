package main

import (
	"fmt"
	"net/http"

	"github.com/wertylop5/yoreo/util"
)

func main() {
	host := ""
	port := 8000
	server := fmt.Sprintf("%s:%d", host, port)
	fileserv := http.FileServer(http.FileSystem(http.Dir("./dist")))

	http.Handle("/", http.StripPrefix("/", fileserv))
	http.HandleFunc("/create-routine", util.CreateRoutine)
	http.HandleFunc("/load-routine", util.LoadRoutine)
	http.HandleFunc("/save-routine", util.SaveRoutine)

	fmt.Printf("listening on %s\n", server)
	http.ListenAndServe(server, nil)
}
