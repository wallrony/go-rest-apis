package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	FirstName string   `json:"first_name,omitempty"`
	LastName  string   `json:"last_name,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

// Main Function
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/contatos", GetPeople).Methods("GET")
	router.HandleFunc("/contatos/add", AddPerson).Methods("POST")
	router.HandleFunc("/contatos/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/contatos/{id}", DeletePerson).Methods("DELETE")

	people = append(people, Person{
		ID: "1", FirstName: "Wallisson", LastName: "Rony", Address: &Address{
			City:  "Arapiraca",
			State: "Alagoas",
		},
	})

	print("Running a server in port 8000")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetPeople(res http.ResponseWriter, req *http.Request) {
	json.NewEncoder(res).Encode(people)
}

func GetPerson(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	var person Person

	for _, item := range people {
		if item.ID == params["id"] {
			person = item
			break
		}
	}

	json.NewEncoder(res).Encode(person)
}

func AddPerson(res http.ResponseWriter, req *http.Request) {
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = fmt.Sprint(len(people) + 1)
	people = append(people, person)

	print(people)

	json.NewEncoder(res).Encode(people)
}

func DeletePerson(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)

			json.NewEncoder(res).Encode(people)

			break
		}
	}

}


