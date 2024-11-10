package database

import (
	"gin-bookstore-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBClient interface {
	DBMigrate() error
}

type Client struct {
	db *gorm.DB
}

func NewClient() (Client, error) {

	dsn := "host=localhost user=postgres password=3141 dbname=Book_Database sslmode=disable"
	dbInfo, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	client := Client{db: dbInfo}

	return client, nil
}

func (c Client) DBMigrate() error {

	err := c.db.AutoMigrate(&models.Book{})

	if err != nil {
		return err
	}

	return nil
}

func (c Client) CloseDBConnection() {
	db, err := c.db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	db.Close()
}
