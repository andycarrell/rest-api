package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	Append(Person{ID: "1", Firstname: "Steve", Lastname: "McStevenson", Address: &Address{City: "City X", State: "State X"}})
	Append(Person{ID: "2", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City Y", State: "State Y"}})
	Append(Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})
}

func main() {
	router := mux.NewRouter()
	routes := InitialiseRoutes()

	for _, route := range routes {
		router.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}

	log.Fatal(http.ListenAndServe(":8000", router))
}
