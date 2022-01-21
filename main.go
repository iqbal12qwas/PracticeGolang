package main

import (
	"practice/db"
	"practice/services"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db.Connect()
	// db.MigratePeople(&entity.People{})
	// db.MigrateFile(&entity.File{})
	services.CreateRouter()
	services.InitializeRoute()
	services.ServerStart()
}
