package parser

import (
	commonobjects "backend/internal/commonObjects"
	commontocsv "backend/internal/parsersFromCommon/commonToCsv"
	csvtocommon "backend/internal/parsersToCommon/csvToCommonParser"
	jsontocommon "backend/internal/parsersToCommon/jsonToCommonParser"
	"log"
)

/*
The main parsing function of the loop designed to
take ParseMultipartForm r.Form contents and parse
from the frontend.
*/
func MainParserLoop(incomingParsingObject commonobjects.ParsingObject) commonobjects.ParsingObject {
	//fmt.Println("Beginning Parsing...")
	//fmt.Println(incomingParsingObject.DataIn)

	// This is what we need all data to become.
	cf := commonobjects.CommonFormat{}
	//var interData map[string]interface{}

	switch incomingParsingObject.DataIn {
	case "JSON":
		log.Println("JSON detected IN")
		cf = jsontocommon.JsonToCommonParser(incomingParsingObject)
		//fmt.Printf("%v", cf.Keys)
	case "CSV":
		log.Println("CSV detected IN")
		csvtocommon.CsvToCommonParser(incomingParsingObject)
		//case 3:
		//    fmt.Println("three")
	}

	switch incomingParsingObject.DataOut {
	case "CSV":
		log.Println("CSV detected OUT!")
		incomingParsingObject.FileOut = commontocsv.CommonToCsvParser(cf)
	}

	//fmt.Printf("%v", incomingParsingObject.FileOut)

	return incomingParsingObject

}
