package api

import (
	//encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"backend/internal/parser"
	"backend/internal/commonObjects"
)

// We have to handle cors options manually here.
// Docs at: https://flaviocopes.com/golang-enable-cors/
func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func errCheck(err error){
	if err != nil {
		fmt.Println(err)
	}
}

// Home page
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received data!")
	setupResponse(&w, r)
	// IF we have a post deal here:
	if (*r).Method == "POST" {
		// 4GB set for MAX Memory size
		r.ParseMultipartForm(4294967296)
		//fmt.Println(r.Form)
		dataOut := r.Form.Get("toFormData")
		dataIn := r.Form.Get("dataIn")

		file, _, err := r.FormFile("fromFiles")
		errCheck(err)
		//fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		//fmt.Println(file)
		fileBytes, err := ioutil.ReadAll(file)
		errCheck(err)
		//fmt.Println(fileBytes)
		parserStruct := commonobjects.ParsingObject {
			DataIn: dataIn,
			DataOut: dataOut,
			FileIn: fileBytes,
		}
		parserStruct = parser.MainParserLoop(parserStruct)



	}
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
