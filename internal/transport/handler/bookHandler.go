package handler

import (
	"crud/internal/core/interface/service"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

type handlerBook struct {
	Title   string `json:"title"`
	Author  string `json:"author"`
	Body    string `json:"body"`
	Release string `json:"release"`
}

func GetBook(service service.BookService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Query("id")

		numberId, err := strconv.Atoi(id)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"message": "неверно передан id книжки"})

			return
		}

		book, err := service.GetBook(c.Request.Context(), numberId)

		if err != nil {
			slog.Error(err.Error())

			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "ошибка получения книжки"})

			return

		}

		c.JSON(http.StatusOK, handlerBook(book))

	}
}

func SearchBooks(service service.BookService) gin.HandlerFunc {
	return func(c *gin.Context) {
		title := c.Query("title")
		author := c.Query("author")
		release := c.Query("release")

		books, err := service.SearchBooks(c.Request.Context(), title, author, release)
		if err != nil {
			slog.Error(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "error retrieving books"})
			return
		}

		c.JSON(http.StatusOK, books)
	}
}

func GetAllBooks(service service.BookService) gin.HandlerFunc {
	return func(c *gin.Context) {
		books, err := service.GetAllBooks(c.Request.Context())
		if err != nil {
			slog.Error(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "ошибка получения списка книг"})
			return
		}
		c.JSON(http.StatusOK, books)
	}
}
