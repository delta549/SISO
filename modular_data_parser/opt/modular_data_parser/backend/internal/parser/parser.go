package parser

import (
	"fmt"
	"backend/internal/jsonParser"
	"backend/internal/commonObjects"
)


/*
The main parsing function of the loop designed to
take ParseMultipartForm r.Form contents and parse
from the frontend.
*/
func MainParserLoop(incomingParsingObject commonobjects.ParsingObject) {
	fmt.Println("Beginning Parsing...")
	//fmt.Println(incomingParsingObject.DataIn)
	
	// This is what we need all data to become.
	var interData map[string]interface{}

    switch incomingParsingObject.DataIn {
    case "JSON":
        fmt.Println("JSON detected!")
		interData = jsonparser.JsonParser(incomingParsingObject)
    //case 2:
    //    fmt.Println("two")
    //case 3:
    //    fmt.Println("three")
    }

	// Parsing json beginning here.
	//fmt.Println(incomingParsingObject.FileIn) // print the content as 'bytes'
    //str := string(incomingParsingObject.FileIn) // convert content to a 'string'
    //fmt.Println(str) // print the content as a 'string'


	fmt.Println(interData)


	//if err := json.Unmarshal(incomingParsingObject.FileIn, &f); err != nil {
	//	panic(err)
	//}


	// Make new temp id for file as a string

	//id := uuid.New()
    //fmt.Println(id.String())

	// write out to file and read in contents 






}