package main

import (
	"fmt"
	"backend/api"
)


func main(){
	fmt.Println("Starting Parser....")

	api.MainApiLoop()

	fmt.Println("Finished.....")

}