package handlers

import (
	"book-management-sqlite/db"
	"book-management-sqlite/models"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBookByID(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	err := db.DB.QueryRow("SELECT id, title, author, isbn, price FROM books WHERE id = ?", id).
		Scan(&book.ID, &book.Title, &book.Author, &book.ISBN, &book.Price)

	if err != nil {
		if err == sql.ErrNoRows {
			c.HTML(http.StatusNotFound, "error.html", gin.H{
				"error": "图书未找到",
			})
		} else {
			c.HTML(http.StatusInternalServerError, "error.html", gin.H{
				"error": "服务器内部错误",
			})
		}
		return
	}

	c.HTML(http.StatusOK, "book_detail.html", gin.H{
		"book": book,
	})
}

func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := db.DB.Exec(
		"INSERT INTO books (title, author, isbn, price) VALUES (?, ?, ?, ?)",
		book.Title, book.Author, book.ISBN, book.Price,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, _ := result.LastInsertId()
	book.ID = int(id)
	c.JSON(http.StatusCreated, book)
}
