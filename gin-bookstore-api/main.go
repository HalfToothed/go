package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string
	Author string
	Read   bool
}

func main() {

	dsn := "host=localhost user=postgres password=3141 dbname=test sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Book{})

	// Create
	db.Create(&Book{Title: "The Hobbit", Author: "J.R.R Tolkien", Read: true})

}
