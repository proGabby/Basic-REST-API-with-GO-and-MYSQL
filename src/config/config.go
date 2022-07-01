package config

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	// "github.com/joho/godotenv"
	env "github.com/willie/BookstoreAPI/cmd"
)

var (
	//using GORM
	//creating an instance for the current db connection
	db *gorm.DB
)

func Connect() {
	//get the db config string from the env key
	dbConfigString := env.GetEnvVariable("DB_CONFIG")
	//initializing db connection "user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", dbConfigString)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	// defer d.Close()
	db = d
}

//function to fetch the currenct db instance
func GetDB() *gorm.DB {
	return db
}
