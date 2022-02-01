package main

import (
	"backend/api"
	"log"
)

func main() {
	log.Println("Starting Parser....")
	apiPort := "3000"
	api.MainApiLoop(apiPort)

	log.Println("Finished.....")

}
