package csvtocommon

import (
	commonobjects "backend/internal/commonObjects"
	"fmt"
	"log"
	"strings"
)


/*
type CommonFormat struct {
	Keys      []string
	KeyValues []interface{}
}
*/

// Hello new?
func CsvToCommonParser(incomingParsingObject commonobjects.ParsingObject) (cf commonobjects.CommonFormat){
	log.Println("Starting CSV parser...")
	// Unmarshall json unknown json string.
	//var input []interface{}

	fmt.Println(string(incomingParsingObject.FileIn))

	stringArray := strings.Split(string(incomingParsingObject.FileIn), "\n")

	fmt.Println(stringArray[0])
	fmt.Println(stringArray[1])
	fmt.Println(stringArray[2])
	//extractKeys()

	// Example of what we have to turn the CSV into programatically.
	cf.Keys = []string{"someKey1", "someKey2"}
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

	return cf




}