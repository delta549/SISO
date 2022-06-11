package parser

import (
	commonobjects "backend/internal/commonObjects"
	filterkeys "backend/internal/filterApplication"
	commontocsv "backend/internal/parsersFromCommon/commonToCsv"
	commontojson "backend/internal/parsersFromCommon/commonToJson"
	commontotsv "backend/internal/parsersFromCommon/commonToTsv"
	csvtocommon "backend/internal/parsersToCommon/csvToCommonParser"
	jsontocommon "backend/internal/parsersToCommon/jsonToCommonParser"
	tsvtocommon "backend/internal/parsersToCommon/tsvToCommonParser"

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
		cf = csvtocommon.CsvToCommonParser(incomingParsingObject)
		//case 3:
		//    fmt.Println("three")
		case "TSV":
		log.Println("TSV detected IN")
		cf = tsvtocommon.TSVToCommonParser(incomingParsingObject)
	}

	if (len(incomingParsingObject.FilterFile) > 0) {
		cf = filterkeys.FilterMain(incomingParsingObject, cf)
	}

	switch incomingParsingObject.DataOut {
	case "CSV":
		log.Println("CSV detected OUT!")
		incomingParsingObject.FileOut = commontocsv.CommonToCsvParser(cf)
	case "JSON":
		log.Println("JSON detected OUT!")
		incomingParsingObject.FileOut = commontojson.CommonToJsonParser(cf)

		//fmt.Println(string(incomingParsingObject.FileOut))
	case "TSV":
		log.Println("TSV detected OUT!")
		incomingParsingObject.FileOut = commontotsv.CommonToTabParser(cf)
	}



	//fmt.Printf("%v", incomingParsingObject.FileOut)

	return incomingParsingObject

}
