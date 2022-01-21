package db

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Connector *gorm.DB
var Connector2 *gorm.DB

func Connect() error {
	config1 :=
		Config{
			ServerName: "localhost:3306",
			User:       "root",
			Password:   "",
			DB:         "user_db",
		}

	config2 :=
		Config{
			ServerName: "localhost:3306",
			User:       "root",
			Password:   "",
			DB:         "practice_java",
		}

	connectionString1 := GetConnectionString(config1)
	connectionString2 := GetConnectionString(config2)

	var err error
	Connector, err = gorm.Open("mysql", connectionString1)
	if err != nil {
		log.Println(err)
	}

	Connector2, err = gorm.Open("mysql", connectionString2)
	if err != nil {
		log.Println(err)
	}

	log.Println("Connection was successful!!")
	return nil
}
