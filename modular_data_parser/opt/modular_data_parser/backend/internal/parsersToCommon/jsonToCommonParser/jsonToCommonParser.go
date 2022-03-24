package jsontocommon

import (
	commonobjects "backend/internal/commonObjects"
	"encoding/json"
	"fmt"
	"reflect"
)

/* JsonParser designed to parse json into a
usable format for conversion into the user
selected format
*/

func extractKeyValues(keyValuesMap map[string]interface{}, keys []string) []string {
	fmt.Println("Extracting Keys")
	for k, _ := range keyValuesMap {
		keys = append(keys, k)
	}
	return keys
}

func extractArrayKeyValues(arrayKeyValues []interface{}) []string {
	keys := []string{}
	if reflect.TypeOf(arrayKeyValues[0]).Kind() == reflect.Map {
		newInput := arrayKeyValues[0].(map[string]interface{})
		keys = extractKeyValues(newInput, keys)
	}
	return keys
}

// Returns a CommonFormat object:
func JsonToCommonParser(incomingParsingObject commonobjects.ParsingObject) (cf commonobjects.CommonFormat) {

	fmt.Println("Starting json file parser...")
	// Unmarshall json unknown json string.
	var input []interface{}
	//cf := commonobjects.CommonFormat{}
	json.Unmarshal(incomingParsingObject.FileIn, &input)
	// Extract known JSON Array.
	cf.Keys = extractArrayKeyValues(input)

	cf.KeyValues = input

	//fmt.Println(cf.Keys)
	//fmt.Printf("%v", cf.KeyValues)

	return cf
}
