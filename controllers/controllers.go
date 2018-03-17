package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/andycarrell/rest-api/data"
	"github.com/gorilla/mux"
)

// Route represents a callable route.
type Route struct {
	Path    string
	Method  string
	Handler func(w http.ResponseWriter, r *http.Request)
}

func getPeople(w http.ResponseWriter, r *http.Request) { json.NewEncoder(w).Encode(data.Get()) }
func getPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range data.Get() {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&data.Person{})
}
func createPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person data.Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	data.Append(person)
	json.NewEncoder(w).Encode(data.Get())
}
func deletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	data.Remove(params["id"])
	json.NewEncoder(w).Encode(data.Get())
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
