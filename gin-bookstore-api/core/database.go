package core

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Client() *gorm.DB {

	dsn := "host=localhost user=postgres password=3141 dbname=test sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return db
}
