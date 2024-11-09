package main

import (
	"gin-bookstore-api/core"
	"gin-bookstore-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var dbClient *gorm.DB

func main() {

	dbClient = core.Client()

	// Migrate the schema
	dbClient.AutoMigrate(&models.Book{})

	router := gin.Default()
	router.POST("/book", postBook)
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
