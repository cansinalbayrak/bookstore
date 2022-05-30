package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/santosh/gingo/db"
	"github.com/santosh/gingo/models"
)

func GetBooks(c *gin.Context) {
	c.JSON(http.StatusOK, db.Books)
}
func PostBook(c *gin.Context) {
	var newBook models.Book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	db.Books = append(db.Books, newBook)
	c.JSON(http.StatusCreated, newBook)
}
func GetBookByISBN(c *gin.Context) {
	isbn := c.Param("isbn")
	for _, a := range db.Books {
		if a.ISBN == isbn {
			c.JSON(http.StatusOK, a)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
}
