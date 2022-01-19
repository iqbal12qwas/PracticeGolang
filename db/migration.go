package db

import (
	"log"
	"practice/entity"
)

func MigratePeople(table *entity.People) {
	Connector.AutoMigrate(&table)
	log.Println("Table People migrated")
}

func MigrateFile(table *entity.File) {
	Connector.AutoMigrate(&table)
	log.Println("Table File migrated")
}
