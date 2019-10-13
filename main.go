package main

import (
	//"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

const fakeApiURL = "https://jsonplaceholder.typicode.com/todos/1"

func root(w http.ResponseWriter, r *http.Request) {
	var result interface{}

	resp, err := http.Get(fakeApiURL)

	if err != nil {
		json.NewEncoder(w).Encode("Error on GET users")
		return
	}

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&result)
	json.NewEncoder(w).Encode(result)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", root).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}