package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vegitobluefan/LibraryAPI/models"
	"github.com/vegitobluefan/LibraryAPI/utils"
)

type BookRequest struct {
	ID string `json:"id"`
}

func GetBooks(c *gin.Context) {
	author := c.Query("author")
	if author == "" {
		c.IndentedJSON(http.StatusOK, models.Books)
		return
	}

	filtered := models.FilterBooksByAuthor(author)
	if len(filtered) == 0 {
		utils.RespondError(c, http.StatusNotFound, "Книги с таким автором не найдены")
		return
	}

	c.IndentedJSON(http.StatusOK, filtered)
}

func GetBookByID(c *gin.Context) {
	id := c.Param("id")
	book, found := models.GetBookByID(id)
	if !found {
		utils.RespondError(c, http.StatusNotFound, "Книга не найдена")
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {
	var newBook models.Book
	if err := c.BindJSON(&newBook); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Некорректный JSON")
		return
	}

	if newBook.ID == "" || newBook.Title == "" || newBook.Author == "" {
		utils.RespondError(c, http.StatusBadRequest, "Все поля книги обязательны")
		return
	}

	models.Books = append(models.Books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func CheckoutBook(c *gin.Context) {
	var req BookRequest
	if err := c.BindJSON(&req); err != nil || req.ID == "" {
		utils.RespondError(c, http.StatusBadRequest, "ID книги не был передан")
		return
	}

	book, found := models.GetBookByID(req.ID)
	if !found {
		utils.RespondError(c, http.StatusNotFound, "Книга не найдена")
		return
	}

	if book.Quantity <= 0 {
		utils.RespondError(c, http.StatusBadRequest, "Книги нет в наличии")
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

func ReturnBook(c *gin.Context) {
	var req BookRequest
	if err := c.BindJSON(&req); err != nil || req.ID == "" {
		utils.RespondError(c, http.StatusBadRequest, "ID книги не был передан")
		return
	}

	book, found := models.GetBookByID(req.ID)
	if !found {
		utils.RespondError(c, http.StatusNotFound, "Книга не найдена")
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}
