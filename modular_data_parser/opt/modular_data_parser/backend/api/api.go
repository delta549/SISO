package api

import (
	"fmt"
	"net/http"
	"log"
)

// Home page 
func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}
// Request handler
func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":3000", nil))
}

// API loop for requests is here:
func MainApiLoop(){
	fmt.Println("Starting API....")
	handleRequests()
}