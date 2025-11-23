package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type config struct{
	PublicHost string
	Port string 
	DbUser  string
	DbPassword string 
	DbAddress string 
	DbName string 
}

var Envs = intconfig()

func intconfig() config {

godotenv.Load()

return config{
    PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
    Port:       getEnv("PORT", "8080"),
    DbUser:     getEnv("DB_USER", "root"),
    DbPassword: getEnv("DB_PASSWORD", "mypassword"),
    DbAddress:  fmt.Sprintf("%s:%s",
                  getEnv("DB_HOST", "127.0.0.1"),
                  getEnv("DB_PORT", "3306")),
    DbName:     getEnv("DB_NAME", "ecom"),
}

}

func getEnv(key, fallback string ) string {
	if value , ok := os.LookupEnv(key); ok {
		return value 
	} 
	return fallback
}


