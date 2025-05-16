package controllers

import (
	"gin-bookstore-api/core/database"

	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	gin *gin.Engine
	db  database.Client
}

func NewServer(db database.Client) *Server {

	server := &Server{

		gin: gin.Default(),
		db:  db,
	}
	server.endpoints()

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
