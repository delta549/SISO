package main

import (
	"backend/api"
	"log"
)

func main() {
	log.Println("Starting Parser....")
	ipAddress := "0.0.0.0"
	apiPort := "8000"
	api.MainApiLoop(apiPort, ipAddress)

	log.Println("Finished.....")

}
