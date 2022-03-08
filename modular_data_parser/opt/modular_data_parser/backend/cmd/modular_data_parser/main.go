package main

import (
	"backend/api"
	"log"
)

func main() {
	log.Println("Starting Parser....")
	apiPort := "8000"
	api.MainApiLoop(apiPort)

	log.Println("Finished.....")

}
