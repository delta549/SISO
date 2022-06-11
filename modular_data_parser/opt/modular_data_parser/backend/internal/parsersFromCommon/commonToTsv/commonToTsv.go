package commontotsv

import (
	"fmt"
	"log"
	//"os"
	commonobjects "backend/internal/commonObjects"
	"strings"
)

func writeValues(cf commonobjects.CommonFormat, overallString string, delim string) string {
	for _, v := range cf.KeyValues{
		line := ""
		switch inter := v.(type){
		case map[string]interface{}:
			for _, key := range cf.Keys{
				line = line + fmt.Sprintf("%v,", inter[key])
			}
			line = strings.TrimRight(line, delim)
			line = line + "\n"
			overallString = overallString + line
		}
	}
	return overallString
}

func CommonToTabParser(cf commonobjects.CommonFormat) ([]byte) {


	//cf.Keys 
	//exampleStringArray := []string{"vegetable", "fruit", "rank"}
	
	//cf.KeyValues
	/*
	exampleKeyValues := []interface{}{
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

	// make file:
	//f, _ := os.Create("output.csv")

	log.Println("Starting CSV parser OUT")

	//fmt.Println(cf.Keys)
	delim := "\t"

	// Make overall string to write.
	overallString := ""

	// Write headers
	singleKey := strings.Join(cf.Keys, delim)
	singleKey = singleKey + "\n"
	overallString = singleKey

	// Write Values
	overallString = writeValues(cf, overallString, delim)

	// Turn into bytes:
	overallStringBytes := []byte(overallString)

	// Test by writing out.
	//f.Write(overallStringBytes)

	//fmt.Printf("%v", overallString)

	log.Println("TAB parser complete")

	return overallStringBytes

}