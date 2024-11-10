package main

import (
	"gin-bookstore-api/core/database"

	"log"

	"github.com/gin-gonic/gin"
)

var dbClient database.DBClient

func main() {

	db, err := database.NewClient()

	if err != nil {
		panic("database Connection failed")
	}

	err = db.DBMigrate()
	if err != nil {
		log.Fatal("Database Migration Failed!")
		return
	}

	dbClient = db

	router := gin.Default()
	// router.GET("/", getAllBooks)
	// router.GET("/book/:id", getBook)
	// router.POST("/book", postBook)
	// router.DELETE("/book/delete/:id", deleteBook)
	router.Run(":8080")
}

// func postBook(c *gin.Context) {
// 	var newBook models.Book

// 	if err := c.BindJSON(&newBook); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Json"})
// 		return
// 	}

// 	if result := dbClient.db.Create(&newBook); result.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
// 		return
// 	}

// 	c.IndentedJSON(http.StatusCreated, newBook)
// }

// func getBook(c *gin.Context) {
// 	id := c.Param("id")

// 	var book models.Book
// 	result := dbClient.First(&book, id)

// 	if result.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
// 		return
// 	}

// 	c.IndentedJSON(http.StatusOK, book)
// }

// func deleteBook(c *gin.Context) {
// 	id := c.Param("id")

// 	var book models.Book
// 	if result := dbClient.Delete(&book, id); result.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
// 		return
// 	}

// 	c.IndentedJSON(http.StatusOK, book)
// }

// func getAllBooks(c *gin.Context) {

// 	var books []models.Book

// 	result := dbClient.Find(&books)

// 	if result.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
// 		return
// 	}

// 	c.IndentedJSON(http.StatusOK, books)
// }
