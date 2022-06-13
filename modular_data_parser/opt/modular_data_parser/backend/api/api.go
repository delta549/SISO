package api

import (
	//"encoding/json"
	commonobjects "backend/internal/commonObjects"
	"backend/internal/parser"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
)

// We have to handle cors options manually here.
// Docs at: https://flaviocopes.com/golang-enable-cors/
func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	//(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	//(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func errCheck(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

// Home page
func homePage(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	// IF we have a post deal here:
	if (*r).Method == "POST" {
		filterFileBytes := []byte{}
		// 4GB set for MAX Memory size
		r.ParseMultipartForm(4294967296)
		dataOut := r.Form.Get("toFormData")
		dataIn := r.Form.Get("dataIn")

		// To print multipartForm file key/values!
		//fmt.Println(r.MultipartForm)

		fromFile, _, err := r.FormFile("fromFiles")
		errCheck(err)
		filterFile, _, filterFileErr := r.FormFile("filterFiles")
		errCheck(err)
		//fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		//fmt.Println(file)
		fromFileBytes, err := ioutil.ReadAll(fromFile)
		errCheck(err)

		// Set filterFileBytes to read if the err is nil
		if filterFileErr == nil {
			filterFileBytes, _ = ioutil.ReadAll(filterFile)
		}


		parserStruct := commonobjects.ParsingObject{
			DataIn:  dataIn,
			DataOut: dataOut,
			FileIn:  fromFileBytes,
			FilterFile: filterFileBytes,
		}
		parserStruct = parser.MainParserLoop(parserStruct)

		// Create Response fto send back to user:
		if dataOut == "JSON"{
			w.Header().Set("Content-Type", "application/json")
		}else{
			w.Header().Set("Content-Type", "application/octet-stream")
		}
		w.Write(parserStruct.FileOut)
	}
}

func test(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	type responseDataStruct struct {
		Response string
	}

	reponseData := responseDataStruct {
		Response: "Success",
	}

	jData, _ := json.Marshal(reponseData)

	w.Write(jData)
}

// Request handler
func handleRequests(apiPort string, ipAddress string) {
	http.HandleFunc("/testendpoint", test)
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", ipAddress ,apiPort), nil))
	//http.ListenAndServe("127.0.0.1:8000", nil)

}

// API loop for requests is here:
func MainApiLoop(apiPort string, ipAddress string) {
	log.Printf("Starting API on port: %s\n", apiPort)
	handleRequests(apiPort, ipAddress)
}
