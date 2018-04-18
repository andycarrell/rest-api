package controllers

import (
	"encoding/json"
	"fmt"
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
	json.NewEncoder(w).Encode(data.GetByID(params["id"]))
}
func createPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]
	if data.GetByID(id).ID == id {
		http.Error(w, fmt.Sprintf("Entity for ID: %s already exists", id), 400)
		return
	}
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
func updatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]
	if data.GetByID(id).ID != id {
		http.Error(w, fmt.Sprintf("Entity for ID: %s doesn't exist", id), 400)
		return
	}
	var person data.Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = id
	data.Replace(id, person)
	json.NewEncoder(w).Encode(data.Get())
}

// InitialiseRoutes returns an array of routes.
func InitialiseRoutes() []Route {
	var routes []Route
	var add = func(r Route) { routes = append(routes, r) }

	add(Route{Path: "/", Method: "GET", Handler: func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("ok")

		headers := w.Header()
		headers.Add("Access-Control-Allow-Origin", "*")

		json.NewEncoder(w).Encode("ok")
	}})
	add(Route{Path: "/error", Method: "GET", Handler: func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("ok")

		headers := w.Header()
		headers.Add("Access-Control-Allow-Origin", "*")

		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}})
	add(Route{Path: "/people", Method: "GET", Handler: getPeople})
	add(Route{Path: "/people/{id}", Method: "GET", Handler: getPerson})
	add(Route{Path: "/people/{id}", Method: "POST", Handler: createPerson})
	add(Route{Path: "/people/{id}", Method: "DELETE", Handler: deletePerson})
	add(Route{Path: "/people/{id}", Method: "PUT", Handler: updatePerson})

	return routes
}
