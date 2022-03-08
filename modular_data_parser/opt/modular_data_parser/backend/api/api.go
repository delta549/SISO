package api

import (
	"fmt"
	"log"
	"net/http"
)

// Home page
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")

	fmt.Println(r.Header)

	fmt.Println(r.Body)

	log.Println("Endpoint Hit: homePage")
}

// Request handler
func handleRequests(apiPort string) {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", apiPort), nil))
}

// API loop for requests is here:
func MainApiLoop(apiPort string) {
	log.Printf("Starting API on port: %s\n", apiPort)
	handleRequests(apiPort)
}
