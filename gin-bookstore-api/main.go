package main

import (
	"gin-bookstore-api/core/database"

	"gin-bookstore-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var dbClient *gorm.DB

func main() {

	dbClient = database.Client()

	// Migrate the schema
	dbClient.AutoMigrate(&models.Book{})

	router := gin.Default()
	router.GET("/", getAllBooks)
	router.GET("/book/:id", getBook)
	router.POST("/book", postBook)
	router.DELETE("/book/delete/:id", deleteBook)
	router.Run(":8080")
}

func postBook(c *gin.Context) {
	var newBook models.Book

	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Json"})
		return
	}

	if result := dbClient.Create(&newBook); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newBook)
}

func getBook(c *gin.Context) {
	id := c.Param("id")

	var book models.Book
	result := dbClient.First(&book, id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func deleteBook(c *gin.Context) {
	id := c.Param("id")

	var book models.Book
	if result := dbClient.Delete(&book, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func getAllBooks(c *gin.Context) {

	var books []models.Book

	result := dbClient.Find(&books)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, books)
}
