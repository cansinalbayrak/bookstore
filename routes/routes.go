package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/santosh/gingo/handlers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/books", handlers.GetBooks)
	router.GET("/books/:isbn", handlers.GetBookByISBN)
	router.POST("/books", handlers.PostBook)
	return router
}
