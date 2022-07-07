package commontojson

import (
	"log"
	commonobjects "backend/internal/commonObjects"
	"encoding/json"
	//"os"
)

// Common object to JSON!
func CommonToJsonParser(cf commonobjects.CommonFormat) ([]byte){
	log.Println("Starting JSON parser OUT")

	marshedBytes, _ := json.MarshalIndent(cf.KeyValues, "", " ")

	//f, _ := os.Create("output.json")


	//f.Write(marshedBytes)

	return marshedBytes
}
