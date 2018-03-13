package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Route represents a callable route.
type Route struct {
	Path    string
	Method  string
	Handler func(w http.ResponseWriter, r *http.Request)
}

func getPeople(w http.ResponseWriter, r *http.Request) { json.NewEncoder(w).Encode(Get()) }
func getPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range Get() {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}
func createPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	Append(person)
	json.NewEncoder(w).Encode(Get())
}
func deletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	Remove(params["id"])
	json.NewEncoder(w).Encode(Get())
}

// InitialiseRoutes returns an array of routes.
func InitialiseRoutes() []Route {
	var routes []Route
	var add = func(r Route) { routes = append(routes, r) }

	add(Route{Path: "/", Method: "GET", Handler: func(w http.ResponseWriter, r *http.Request) { json.NewEncoder(w).Encode("ok") }})
	add(Route{Path: "/people", Method: "GET", Handler: getPeople})
	add(Route{Path: "/people/{id}", Method: "GET", Handler: getPerson})
	add(Route{Path: "/people/{id}", Method: "POST", Handler: createPerson})
	add(Route{Path: "/people/{id}", Method: "DELETE", Handler: deletePerson})

	return routes
}
