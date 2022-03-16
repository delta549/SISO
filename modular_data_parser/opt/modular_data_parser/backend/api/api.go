package api

import (
	//encoding/json"
	"fmt"
	"log"
	"net/http"
)

// We have to handle cors options manually here.
// Docs at: https://flaviocopes.com/golang-enable-cors/
func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

// Home page
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received data!")
	setupResponse(&w, r)
	// IF we have a post deal here:
	if (*r).Method == "POST" {
		// 4GB set for MAX Memory size
		r.ParseMultipartForm(4294967296)
		fmt.Println(r.Form)
	}
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
