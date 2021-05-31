package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

//Type definitions go here
type User struct {
	ID       int    `json:"id"`
	Username string `json: "username"`
	Email    string `json: "email"`
}

type Users struct {
	Users []User
}

type Error struct {
	ID          int    `json:"id"`
	Description string `json: "description"`
}

type Errors struct {
	Errors []Error
}

//Global variables
var users Users   //To keep JSON user data
var errors Errors //To keep error messages

//Main function
func main() { //Main function
	getErrorMessages()
	getUsersFromFile()
	fmt.Printf("Serving HTTP requests.\n")
	handleRequests()
}

//File access functions go here
func getUsersFromFile() { //Get user data from user.json
	jsonFile, err := os.Open("Server/user.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &users)
}

func getErrorMessages() { //Get API server error messages from error.json
	jsonFile, err := os.Open("Server/error.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &errors)
}

//API request handler
func handleRequests() { //Request handler
	apiRouter := mux.NewRouter().StrictSlash(true) //Create a router using mux to handle arguments
	//List of requests served
	apiRouter.HandleFunc("/", homePage)                    //Home page
	apiRouter.HandleFunc("/getall", getAllUsers)           //Get and serve all user data in user.json
	apiRouter.HandleFunc("/get/{id}", getSingleUser)       //Get and serve single user data by ID
	apiRouter.HandleFunc("/deleteall", removeUserData)     //Remove all user data
	apiRouter.HandleFunc("/delete/{id}", removeSingleUser) //Remove single user's data by ID
	apiRouter.HandleFunc("/commit", commitChanges)         //Flush cache to user.json
	apiRouter.HandleFunc("/discard", discardChanges)       //Discard cache and reload data from user.json
	//Listen for requests on port 10000
	log.Fatal(http.ListenAndServe(":10000", apiRouter))
}

//API endpoints go here
func homePage(w http.ResponseWriter, r *http.Request) { //Homepage
	fmt.Fprintf(w, "Home page")
	fmt.Printf("Enpoint hit: homePage\n")
}

func getAllUsers(w http.ResponseWriter, r *http.Request) { //Get all users on request
	json.NewEncoder(w).Encode(users.Users)
	fmt.Printf("Enpoint hit: getAllUsers\n")
}

func getSingleUser(w http.ResponseWriter, r *http.Request) { //Get single user by ID
	vars := mux.Vars(r)
	key, err := strconv.Atoi(vars["id"])
	if err != nil {
		json.NewEncoder(w).Encode(errors.Errors[0])
	} else {
		served := false
		for _, user := range users.Users {
			if user.ID == key {
				json.NewEncoder(w).Encode(user)
				served = true
			}
		}
		if !served {
			json.NewEncoder(w).Encode(errors.Errors[1])
		}
	}
	fmt.Printf("Enpoint hit: getSingleUser\n")
}

func removeUserData(w http.ResponseWriter, r *http.Request) { //Remove all user data
	//This function needs to be cached and require a commit request to perform. Might need to rewrite to clear cached user data.
	//os.Remove("Server/user.json")
	fmt.Printf("Enpoint hit: removeUserData\n")
}

func removeSingleUser(w http.ResponseWriter, r *http.Request) { //Remove user data by ID
	//Function not yet ready.
	fmt.Printf("Enpoint hit: removeSingleUser\n")
}

func commitChanges(w http.ResponseWriter, r *http.Request) { //Flush cache to user.json
	//Function not yet ready.
	fmt.Printf("Enpoint hit: commitChanges\n")
}

func discardChanges(w http.ResponseWriter, r *http.Request) { //Discard changes to cache, reload data from user.json
	//Function not yet ready
	fmt.Printf("Enpoint hit: discardChanges\n")
}
