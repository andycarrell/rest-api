package main

import (
	"log"
	"net/http"

	"github.com/andycarrell/rest-api/controllers"
	"github.com/andycarrell/rest-api/data"
	"github.com/gorilla/mux"
)

func init() {
	data.Append(data.Person{
		ID:        "1",
		Firstname: "Steve",
		Lastname:  "McStevenson",
		Address: &data.Address{
			City:  "City X",
			State: "State X",
		},
	})
	data.Append(data.Person{
		ID:        "2",
		Firstname: "John",
		Lastname:  "Doe",
		Address: &data.Address{
			City:  "City Y",
			State: "State Y",
		},
	})
	data.Append(
		data.Person{
			ID:        "3",
			Firstname: "Francis",
			Lastname:  "Sunday",
		},
	)
}

func main() {
	router := mux.NewRouter()
	routes := controllers.InitialiseRoutes()

	for _, route := range routes {
		router.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}

	log.Fatal(http.ListenAndServe(":8000", router))
}
