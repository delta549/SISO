package filterkeys

import (
	commonobjects "backend/internal/commonObjects"
	"encoding/json"
	//"fmt"
	"log"
	//"reflect"
)

func filterOutKeys(trueKey string, cf commonobjects.CommonFormat, newKeys []string) []string {
	for _, key := range cf.Keys{
		if key == trueKey{
			newKeys = append(newKeys, key)
		}
	}
	return newKeys
}

func filterOutKeyValues(cf commonobjects.CommonFormat, newKeys []string) []interface{}{
	var newMappedObjectArray []interface{}
	for _, originalObject := range cf.KeyValues{
		newMappedOriginalObject := make(map[string]interface{})
		for _, key := range newKeys{
			mappedOriginalObject := originalObject.(map[string]interface{})
			for k, v := range mappedOriginalObject{
				if key == k{
					var x interface{} = v
					newMappedOriginalObject[key] = x
				}
			}
		}
		newMappedObjectArray = append(newMappedObjectArray, newMappedOriginalObject)
	}

	return newMappedObjectArray
}

func extractKeyValues(keyValuesMap map[string]interface{}, keys []string) []string {
	//fmt.Println("Extracting Keys")
	for k, _ := range keyValuesMap {
		keys = append(keys, k)
	}
	return keys
}

func extractFilterKeyValues(filterKeyValues interface{}, cf commonobjects.CommonFormat,
	newKeys []string, newKeyValues []interface{}) ([]string, []interface{}){
	switch filterKeyValues.(type){
	case map[string]interface{}:
		somethingnew := filterKeyValues.(map[string]interface{})
		for trueKey, v := range somethingnew {
			if v == true {
				newKeys = filterOutKeys(trueKey, cf, newKeys)
			}
		}
	}

	newKeyValues = filterOutKeyValues(cf, newKeys)
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

	cf.KeyValues = newKeyValues

	return cf
}