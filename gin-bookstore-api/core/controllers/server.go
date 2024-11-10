package controllers

import (
	"gin-bookstore-api/core/database"

	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	gin *gin.Engine
	db  database.DBClient
}

func NewServer(db database.DBClient) *Server {

	server := &Server{

		gin: gin.Default(),
		db:  db,
	}

	return server
}

func (s *Server) Start() error {

	err := s.gin.Run(":8080")
	if err != nil {
		log.Fatalf("Server Issue: %s", err)
		return err
	}

	return nil
}
