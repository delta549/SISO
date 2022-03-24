package commontocsv

import (
	"fmt"
	"os"
	"strings"
	commonobjects "backend/internal/commonObjects"
)

func writeValues(cf commonobjects.CommonFormat, f *os.File, overallString string) string {
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
	f, _ := os.Create("output.csv")

	// Make overall string to write.
	overallString := ""

	// Write headers
	singleKey := strings.Join(cf.Keys, ",")
	singleKey = singleKey + "\n"
	overallString = singleKey

	// Write Values
	overallString = writeValues(cf, f, overallString)

	// Turn into bytes:
	overallStringBytes := []byte(overallString)

	f.Write(overallStringBytes)

	fmt.Printf("%v", overallString)

	return overallStringBytes

}
