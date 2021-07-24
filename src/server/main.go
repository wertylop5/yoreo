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

	err := util.InitDb("./temp.db")
	if err != nil {
		panic(fmt.Sprintf("couldn't initialize db: %s\n", err.Error()))
	}
	util.InitTables()

	http.Handle("/", http.StripPrefix("/", fileserv))
	http.HandleFunc("/routine/create", util.CreateRoutine)
	http.HandleFunc("/routine/load", util.LoadRoutine)
	http.HandleFunc("/routine/save", util.SaveRoutine)

	http.HandleFunc("/user/create", util.CreateUser)
	http.HandleFunc("/user/get", util.GetUser)

	fmt.Printf("listening on %s\n", server)
	http.ListenAndServe(server, nil)
}
