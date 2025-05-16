package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:title  binding:"required"`
	Author string `json:author binding:"required"`
	Read   bool   `json:read`
}

type BookParams struct {
	Id     int64  `json:id`
	Title  string `json:title`
	Author string `json:author`
	Read   bool   `json:read`
}

type UpdateBookParams struct {
	Title  string `json:title  binding:"required"`
	Author string `json:author binding:"required"`
	Read   bool   `json:read`
}
