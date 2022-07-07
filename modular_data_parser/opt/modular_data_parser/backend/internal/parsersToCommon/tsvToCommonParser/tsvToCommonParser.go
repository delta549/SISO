package tsvtocommonparser

import (
	commonobjects "backend/internal/commonObjects"
	"fmt"
	//"fmt"
	//"fmt"
	"log"
	"strings"
)


func extractArrayKeyValues(stringArray []string, cf commonobjects.CommonFormat, delim string) []interface{}{
	var keyValues []interface{}
	for i , line := range stringArray {
		fmt.Println(line)
		if i == 0 || line == ""{
			continue
		}
		splitCsvLine := strings.Split(line, delim)
		fmt.Println(splitCsvLine[0])
		currentKeyValueLine := map[string]interface{}{}
		for keypos , key := range cf.Keys {
			currentKeyValueLine[key] = splitCsvLine[keypos]
			//fmt.Println(currentKeyValueLine[key])
		}
		keyValues = append(keyValues, currentKeyValueLine)
	}
	return keyValues
}

// Special attempts to use any delemiter value:
func TSVToCommonParser(incomingParsingObject commonobjects.ParsingObject) (cf commonobjects.CommonFormat){
	log.Println("Starting TAB parser IN")

	delim := "\t"

	stringArray := strings.Split(string(incomingParsingObject.FileIn), "\n")
	
	// headings or keys is always top line in csv files.
	cf.Keys = strings.Split(stringArray[0], delim)
	//fmt.Println(cf.Keys)

	cf.KeyValues = extractArrayKeyValues(stringArray, cf, delim)

	//fmt.Println(stringArray[1])
	//fmt.Println(stringArray[2])
	//extractKeys()

	// Example of what we have to turn the CSV into programatically.
	//cf.Keys = []string{"someKey1", "someKey2"}
	/*
	cf.KeyValues = []interface{}{
		map[string]interface{}{
			"vegetable": "carrot",
			"fruit":     "banana",
			"rank":      1,
		},
		map[string]interface{}{
			"vegetable": "potato",
			"fruit":     "strawberry",
			"rank":      2,
		},
	}
	*/

	return cf




}