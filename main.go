package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vegitobluefan/LibraryAPI/handlers"
)

func main() {
	router := gin.Default()

	router.GET("/books", handlers.GetBooks)
	router.GET("/books/:id", handlers.GetBookByID)
	router.POST("/books", handlers.CreateBook)
	router.POST("/books/checkout", handlers.CheckoutBook)
	router.POST("/books/return", handlers.ReturnBook)

	router.Run("localhost:8080")
}
