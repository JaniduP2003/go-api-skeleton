package main

import (
	"backend/cmd/api"
	"backend/config"
	"backend/db"
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {

	db, err :=db.NewMySQLStorage(mysql.Config{
		User :config.Envs.DbUser,
		Passwd: config.Envs.DbPassword,
		Addr: config.Envs.DbAddress,
		DBName: config.Envs.DbName,
		Net: "tcp",
		AllowNativePasswords: true ,
		ParseTime: true ,

	})
	if err !=nil{
		log.Fatal(err)
	}


	// Application entry point
	server := api.NewAPIserver(":8080" , nil)
	if err := server.Run() ; err!=nil{
		log.Fatal(err)
	}
}
