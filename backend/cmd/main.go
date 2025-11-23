package main

import (
	"backend/cmd/api"
	"backend/config"
	"backend/db"
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {

	db, err :=db.MySqlStorage(mysql.Config{
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

	initStorgage(db)


	// Application entry point
	server := api.NewAPIserver(":8080" , nil)
	if err := server.Run() ; err!=nil{
		log.Fatal(err)
	}
}

func initStorgage(db *sql.DB ){
	err :=db.Ping()
	if err !=nil{
		log.Fatal(err)
	}
	log.Println("DB:successfulyy caonncated !!! damn ")
}
