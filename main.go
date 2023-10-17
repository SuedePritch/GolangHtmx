package main

import (
	"log"
	"net/http"

	_ "github.com/libsql/libsql-client-go/libsql"
)

func main() {
	db = ConnectToDatabase()
	http.HandleFunc("/", Homepage)
	http.HandleFunc("/user/", ToggleEditUser)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
