package main

import (
	"backend/cmd/api"
	"backend/db"
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {

	db,err :=db.MySqlStorage(*mysql.NewConfig()){
		(User "root"
		Passwd "asd",
		Addr "127.0.1:3306",
		DBName "dbdb",
		Net "tcp",
		AllowNativePassword "true",
		ParseTime: "ture", )
	}


	// Application entry point
	server := api.NewAPIserver(":8080" , nil)
	if err := server.Run() ; err!=nil{
		log.Fatal(err)
	}
}
