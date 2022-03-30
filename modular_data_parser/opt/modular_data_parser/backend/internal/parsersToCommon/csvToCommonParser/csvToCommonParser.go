package csvtocommon

import (
	commonobjects "backend/internal/commonObjects"
	//"fmt"
	//"fmt"
	"log"
	"strings"
)

/*
type CommonFormat struct {
	Keys      []string
	KeyValues []interface{}
}
*/

func extractArrayKeyValues(stringArray []string, cf commonobjects.CommonFormat) []interface{}{
	var keyValues []interface{}
	for i , line := range stringArray {
		if i == 0 || line == ""{
			continue
		}
		splitCsvLine := strings.Split(line, ",")
		currentKeyValueLine := map[string]interface{}{}
		for keypos , key := range cf.Keys {
			currentKeyValueLine[key] = splitCsvLine[keypos]
		}
		keyValues = append(keyValues, currentKeyValueLine)
	}
	return keyValues
}

// Hello new?
func CsvToCommonParser(incomingParsingObject commonobjects.ParsingObject) (cf commonobjects.CommonFormat){
	log.Println("Starting CSV parser IN")
	// Unmarshall json unknown json string.
	//var input []interface{}

	//fmt.Println(string(incomingParsingObject.FileIn))
	// Split into seperate newlines
	stringArray := strings.Split(string(incomingParsingObject.FileIn), "\n")
	// headings or keys is always top line in csv files.
	cf.Keys = strings.Split(stringArray[0], ",")

	cf.KeyValues = extractArrayKeyValues(stringArray, cf)

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