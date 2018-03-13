package main

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

// Append p to the people list
func Append(p Person) []Person {
	people = append(people, p)
	return people
}

// Get list of people
func Get() []Person { return people }

// Remove a person
func Remove(id string) []Person {
	for index, item := range Get() {
		if item.ID == id {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	return people
}
