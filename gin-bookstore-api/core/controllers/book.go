package controllers

import (
	"fmt"
	"gin-bookstore-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Server) CreateBook(c *gin.Context) {

	var book models.BookParams

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	addBook, err := s.db.AddBook(c, book)
	if err != nil {
		fmt.Println("CreateErr", err.Error())
	}
	c.JSON(http.StatusOK, addBook)

}

func (s *Server) ListBooks(c *gin.Context) {

	listBooks, err := s.db.ListBooks(c)

	if err != nil {
		fmt.Println("err", err.Error())
	}

	c.JSON(http.StatusOK, listBooks)
}

func (s *Server) GetBookById(c *gin.Context) {

	bookId := c.Param("id")
	id, _ := strconv.ParseInt(bookId, 10, 64)

	book, err := s.db.GetBook(c, id)

	if err != nil {
		fmt.Println("err", err.Error())
	}

	c.JSON(http.StatusOK, book)
}

func (s *Server) UpdateBook(c *gin.Context) {

	var book models.UpdateBookParams

	bookId := c.Param("id")
	id, _ := strconv.ParseInt(bookId, 10, 64)

	if err := c.ShouldBindJSON(&book); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Check the data you are passing."})
		return
	}

	updatedBook, err := s.db.UpdateBook(c, book, id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if updatedBook {
		c.JSON(http.StatusOK, gin.H{"status": true, "message": "Book Updated!"})
		return
	}
}

func (s *Server) DeleteBook(c *gin.Context) {

	bookId := c.Param("id")
	id, _ := strconv.ParseInt(bookId, 10, 64)

	err := s.db.DeleteBook(c, id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Book Deleted!"})
	return

}
