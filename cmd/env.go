package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

//set the env file and get value from an env key
var GetEnvVariable = func(key string) string {
	//load the env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file")
	}

	//lookup for the key
	value, exists := os.LookupEnv(key)

	if !exists {
		log.Fatal("no env variable found for the key")
	}
	//return the value found on the key
	return value
}
