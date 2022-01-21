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

func MigrateBook(table *entity.Book) {
	Connector2.AutoMigrate(&table)
	log.Println("Table Book migrated")
}
