package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var cache map[string]string

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/data", FileLineSearch).Methods("GET")

	port := ":8080"

	log.Println("Starting the server at port ", port)
	log.Println("Starting indexing of data")
	precompute_data()
	err := http.ListenAndServe(port, router)

	if err != nil {
		log.Println("error starting the server at port", port)
		return
	}

}
