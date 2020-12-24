package main

import (
	"log"
	"net/http"

	"./routers"

	f "fmt"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/users", routers.IndexUsers).Methods("GET")
	router.HandleFunc("/users", routers.AddUser).Methods("POST")
	router.HandleFunc("/users/{id}", routers.ShowUser).Methods("GET")
	router.HandleFunc("/users/{id}", routers.DeleteUser).Methods("DELETE")

	f.Println("Server online in port 8000.")

	log.Fatal(http.ListenAndServe(":8000", router))
}
