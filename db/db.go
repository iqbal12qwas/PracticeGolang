package db

import (
	"log"
	"practice/entity"

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

func MigratePeople(table *entity.People) {
	Connector.AutoMigrate(&table)
	log.Println("Table People migrated")
}

func MigrateFile(table *entity.File) {
	Connector.AutoMigrate(&table)
	log.Println("Table File migrated")
}
