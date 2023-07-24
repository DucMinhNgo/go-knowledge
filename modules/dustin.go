package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
)

type MyData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func checkHeath(response http.ResponseWriter, request *http.Request) {
	data := MyData{
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}

	responseJSON, err := json.Marshal(data)
	if err != nil {
		http.Error(response, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	response.Write(responseJSON)
}

func main () {
	fmt.Printf("Dustin")
	router := mux.NewRouter()

	router.HandleFunc("/", checkHeath).Methods("GET")

	fmt.Printf("Golang Rest API Is Running On Port: 5000")

	err := http.ListenAndServe(":5000", router)

	if err != nil {
		panic(err)
	}
}