package main

import (
	"database/sql"
	"net/http"
)

type Service struct {
	db *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{db: db}
}

func (s *Service) CreateProduct(writer http.ResponseWriter, request *http.Request) {

}

func (s *Service) GetProduct(writer http.ResponseWriter, request *http.Request) {

}
