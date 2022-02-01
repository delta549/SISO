package api

import (
	"fmt"
	"log"
	"net/http"
)

// Home page
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

// Request handler
func handleRequests(api string) {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", api), nil))
}

// API loop for requests is here:
func MainApiLoop() {
	fmt.Println("Starting API....")
	apiPort := "3000"
	handleRequests(api)
}
