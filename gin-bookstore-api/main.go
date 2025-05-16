package main

import (
	"gin-bookstore-api/core/controllers"
	"gin-bookstore-api/core/database"

	"log"
)

func main() {

	db, err := database.NewClient()

	if err != nil {
		panic("database Connection failed")
	}

	err = db.DBMigrate()
	if err != nil {
		log.Fatal("Database Migration Failed!")
		return
	}

	service := controllers.NewServer(db)
	log.Fatal(service.Start())
}
