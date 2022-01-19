package db

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//Connector variable used for CRUD operation's
var Connector *gorm.DB

//Connect creates MySQL connection
func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	log.Println("Connection was successful!!")
	return nil
}

func CloseDatabase(connection *gorm.DB) {
	sqldb := connection.DB()
	sqldb.Close()
}

// Init DB
func DBinit() {
	config :=
		Config{
			ServerName: "localhost:3306",
			User:       "root",
			Password:   "",
			DB:         "user_db",
		}

	connectionString := GetConnectionString(config)
	err := Connect(connectionString)
	if err != nil {
		log.Fatal(err)
	}
}
