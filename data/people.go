package data

// Person entity.
type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

// Address entity.
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

// Append p to the people list.
func Append(p Person) []Person {
	people = append(people, p)
	return people
}

// Get list of people.
func Get() []Person { return people }

// GetByID returns person for the given ID, or default.
func GetByID(id string) Person {
	for _, person := range people {
		if person.ID == id {
			return person
		}
	}
	return Person{}
}

// Remove a person.
func Remove(id string) []Person {
	for index, item := range Get() {
		if item.ID == id {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	return people
}

// Replace a person.
func Replace(id string, p Person) []Person {
	for index, item := range Get() {
		if item.ID == id {
			people = append(append(people[:index], p), people[index+1:]...)
			break
		}
	}
	return people
}
