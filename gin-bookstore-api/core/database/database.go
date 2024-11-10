package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBClient interface {
	DBMigrate() error
}

type Client struct {
	db *gorm.DB
}

func NewClient() (DBClient, error) {

	dsn := "host=localhost user=postgres password=3141 dbname=test sslmode=disable TimeZone=Asia/Shanghai"
	dbInfo, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	client := Client{db: dbInfo}

	return client, nil
}

func (c Client) DBMigrate() error {
	return nil
}
