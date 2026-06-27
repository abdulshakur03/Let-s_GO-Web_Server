package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/snippet/create", snippetCreate)
	mux.HandleFunc("/snippet/view", snippetView)
	log.Print("Server Started on: http://localhost:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
