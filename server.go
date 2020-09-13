package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	const port string = ":8000"

	r := mux.NewRouter()

	r.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Up and Running")
	})

	r.HandleFunc("/posts", getPosts).Methods("GET")
	r.HandleFunc("/posts", addPost).Methods("POST")
	log.Fatal(http.ListenAndServe(port, r))
}
