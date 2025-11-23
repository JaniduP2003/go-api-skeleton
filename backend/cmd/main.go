package main

import (
	"backend/cmd/api"
	"log"
)

func main() {
	// Application entry point
	server := api.NewAPIserver(":8080" , nil)
	if err := server.Run() ; err!=nil{
		log.Fatal(err)
	}
}
