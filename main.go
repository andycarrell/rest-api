package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func populate() {
	Append(Person{ID: "1", Firstname: "Steve", Lastname: "McStevenson", Address: &Address{City: "City X", State: "State X"}})
	Append(Person{ID: "2", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City Y", State: "State Y"}})
	Append(Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})
}

func main() {
	populate()
	router := mux.NewRouter()
	router.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) { json.NewEncoder(w).Encode("ok") },
	).Methods("GET")
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
