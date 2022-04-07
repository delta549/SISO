package filterkeys

import (
	commonobjects "backend/internal/commonObjects"
	"encoding/json"
	"fmt"
	"log"
)

func filterOut(trueKey string, cf commonobjects.CommonFormat, newKeys []string) []string{
	for _, key := range cf.Keys{
		if key == trueKey{
			newKeys = append(newKeys, key)
		}
	}
	return newKeys
}


func extractFilterKeyValues(filterKeyValues interface{}, cf commonobjects.CommonFormat,
	newKeys []string, newKeyValues []interface{}) ([]string, []interface{}){
	switch filterKeyValues.(type){
	case map[string]interface{}:
		somethingnew := filterKeyValues.(map[string]interface{})
		for trueKey, v := range somethingnew {
			if v == true {
				newKeys = filterOut(trueKey, cf, newKeys)
			}
			for _, obj := range cf.KeyValues{
				switch obj.(type){
				case map[string]interface{}:
					newobj := obj.(map[string]interface{})
					fmt.Println(trueKey)
					delete(newobj, trueKey)
					newKeyValues = append(newKeyValues, newobj)
				}
			}
		}
	}

	return newKeys, newKeyValues
}


// The main filter method:
func FilterMain(incomingParsingObject commonobjects.ParsingObject, cf commonobjects.CommonFormat) commonobjects.CommonFormat{
	log.Println("Starting Filter")

	// Unmarshall json unknown json string.
	var input interface{}

	var newKeys []string

	var newKeyValues []interface{}

	json.Unmarshal(incomingParsingObject.FilterFile, &input)

	newKeys, newKeyValues = extractFilterKeyValues(input, cf, newKeys, newKeyValues)

	cf.Keys = newKeys

	//cf.KeyValues = newKeyValues

	return cf
}