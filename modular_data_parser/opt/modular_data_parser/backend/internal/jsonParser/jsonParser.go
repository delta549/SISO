package jsonparser

import (
	"fmt"
	"backend/internal/commonObjects"
	"encoding/json"
)


/* JsonParser designed to parse json into a 
usable format for conversion into the user 
selected format
*/

func JsonParser(incomingParsingObject commonobjects.ParsingObject) (returnedInterData map[string]interface{}){

	// Get data into correct format
	// Filter out fields
	// Return for conversion.
	var interData map[string]interface{}
	keys := []string{}
	
	//fmt.Println(keys)
	json.Unmarshal(incomingParsingObject.FileIn, &interData)
	// Walk over all key/value pairs:
	keys, _ = recursiveWalker(interData, keys)

	fmt.Println(keys)


	//fmt.Println(interData)

	return interData
}


func recursiveWalker(interData map[string]interface{}, keys []string) ([]string, []string){
	values := make([]string, 0, len(interData))
	for k, val := range interData {
		fmt.Printf("Key: %v, Value: %T\n" ,k,val)
		keys = append(keys, k)
		switch vv := val.(type) {
		case string:
		  fmt.Printf("Item %q is a string, containing %q\n", k, vv)
		  fmt.Println(keys)
		case float64:
		  fmt.Printf("Looks like item %q is a number, specifically %f\n", k, vv)
		case map[string]interface{}:
			//fmt.Printf("INTERFACEMAP: Key: %v Type: %v\n", k, vv)
			retvalues, _ := recursiveWalker(vv, keys)
            for _, value := range retvalues {
                fmt.Printf("RECURSION: %v\n", value)
            }
		case map[string]string:
		  fmt.Println("String detected!")
		case map[string]float64:
		  fmt.Println("Float detected!")
		case []interface{}:
		  retvalues, _ := recursiveInterfaceWalker(vv, keys)
		  for _, value := range retvalues {
			fmt.Printf("RECURSION: %v\n", value)
		}
		default:
		  fmt.Printf("Not sure what type item %q is, but I think it might be %T\n", k, vv)
		}
	}
	return values, keys
}

func recursiveInterfaceWalker(interData []interface{}, keys []string) ([]string, []string){
	
	values := make([]string, 0, len(interData))
	for k, val := range interData {
		fmt.Printf("RECURSIVE INTERFACE Key: %v, Value: %v%T\n" ,k,val,val)
		switch vv := val.(type) {
		case string:
			fmt.Println("hi")
		  //fmt.Printf("Interface Item %q is a string, containing %q\n", k, vv)
		case float64:
		  fmt.Printf("Looks like item %q is a number, specifically %f\n", k, vv)
		case map[string]interface{}:
			retvalues, _ := recursiveWalker(vv, keys)
            for _, value := range retvalues {
                fmt.Printf("RECURSION: %v\n", value)
            }
		case map[string]string:
		  fmt.Println("String detected!")
		case map[string]float64:
		  fmt.Println("Float detected!")
		case []interface{}:
		  fmt.Printf("INTERFACE: %v\n", vv)
		  retvalues, _ := recursiveInterfaceWalker(vv, keys)
		  for _, value := range retvalues {
			fmt.Printf("RECURSION: %v\n", value)
		}
		default:
		  fmt.Printf("Not sure what type item %q is, but I think it might be %T\n", k, vv)
		}
	}
	return values, keys
}


