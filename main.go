package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	//"errors"
)

type book struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Quantity int `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel", Quantity: 2},
	{ID: "2", Title: "Text 2", Author: "Author 2", Quantity: 4},
	{ID: "3", Title: "Text 3", Author: "Author 3", Quantity: 5},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.Run("localhost:8080")
}