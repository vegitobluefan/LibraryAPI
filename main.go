package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "Преступление и наказание", Author: "Фёдор Достоевский", Quantity: 3},
	{ID: "2", Title: "Великий Гэтсби", Author: "Ф. Скотт Фицджеральд", Quantity: 5},
	{ID: "3", Title: "Война и мир", Author: "Лев Толстой", Quantity: 6},
	{ID: "4", Title: "Мастер и Маргарита", Author: "Михаил Булгаков", Quantity: 2},
	{ID: "5", Title: "Евгений Онегин", Author: "Александр Пушкин", Quantity: 4},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func getBookFromQuery(c *gin.Context) (*book, bool) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID книги не был передан!"})
		return nil, false
	}

	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Книга не найдена!"})
		return nil, false
	}

	return book, true
}

func bookById(c *gin.Context) {
	id := c.Param(("id"))
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Книга не найдена!"})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func checkoutBook(c *gin.Context) {
	book, ok := getBookFromQuery(c)
	if !ok {
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Книги нет в наличии!"})
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

func returnBook(c *gin.Context) {
	book, ok := getBookFromQuery(c)
	if !ok {
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}

func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("book not found")
}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)
	router.POST("/books", createBook)
	router.PATCH("/checkout", checkoutBook)
	router.PATCH("/return", returnBook)
	router.Run("localhost:8080")
}
