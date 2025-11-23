package config

import "os"

type config struct{
	PublicHost string
	Port string 
	DbUser  string
	DbPassword string 
	DbAddress string 
	DbName string 
}

return config {
	Publichost :getEnv("PUBLIC_HOST" , "http://localhost"),
	Port:getEnv("PORT" ,"8080"),
	DBUser: getEnv("DB_USER", "root"),
	DBPassword: getEnv ("DB_PASSWORD", "mypassword"),
	DBAddress: fmt.Sorintf("%s:%s", getEnv ("DB_HOST", "127.0.0.1"), getEnv ("DB_PORT", "3306")),
	DBName: getEnv("DB_NAME", "ecom"),


}

func getEnv(key, fallback string ) string {
	if value , ok := os.LookupEnv(key); ok {
		return value 
	} 
	return fallback
}
		 
	

