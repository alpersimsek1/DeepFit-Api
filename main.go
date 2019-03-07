package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

// Article - Our struct for all articles
type UserInput struct {
	ID        string `json:"id"`
	Timestamp string `json:"timestamp"`
	Input     string `json:"input"`
	Content   string `json:"content"`
}

type UserInputs []UserInput

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllUserInputs(w http.ResponseWriter, r *http.Request) {
	inputs := UserInputs{
		UserInput{ID: "123", Timestamp: "12:23:22", Input: "cheese", Content: "meal"},
		UserInput{ID: "456", Timestamp: "12:23:25", Input: "chicken", Content: "meal"},
	}
	fmt.Println("Endpoint Hit: returnAllUserInputs")

	json.NewEncoder(w).Encode(inputs)
}

func returnSingleInput(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["ID"]
	fmt.Fprintf(w, "Key: "+key)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", returnAllUserInputs)
	myRouter.HandleFunc("/article/{ID}", returnSingleInput).Methods("GET")
	myRouter.HandleFunc("/article/{ID}", returnSingleInput).Methods("POST")

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	handleRequests()
}