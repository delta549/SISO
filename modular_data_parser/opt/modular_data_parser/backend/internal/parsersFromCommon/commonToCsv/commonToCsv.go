package commontocsv

import (
	"fmt"
	"log"
	//"os"
	commonobjects "backend/internal/commonObjects"
	"strings"
)

func writeValues(cf commonobjects.CommonFormat, overallString string) string {
	for _, v := range cf.KeyValues{
		line := ""
		switch inter := v.(type){
		case map[string]interface{}:
			for _, key := range cf.Keys{
				line = line + fmt.Sprintf("%v,", inter[key])
			}
			line = strings.TrimRight(line, ",")
			line = line + "\n"
			overallString = overallString + line
		}
	}
	return overallString
}

func CommonToCsvParser(cf commonobjects.CommonFormat) ([]byte) {


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

	// Make overall string to write.
	overallString := ""

	// Write headers
	singleKey := strings.Join(cf.Keys, ",")
	singleKey = singleKey + "\n"
	overallString = singleKey

	// Write Values
	overallString = writeValues(cf, overallString)

	// Turn into bytes:
	overallStringBytes := []byte(overallString)

	// Test by writing out.
	//f.Write(overallStringBytes)

	//fmt.Printf("%v", overallString)

	log.Println("CSV parser complete")

	return overallStringBytes

}
