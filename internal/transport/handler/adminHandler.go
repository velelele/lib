package handler

import (
	"crud/internal/core/interface/service"
	"crud/internal/core/model"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func AddBook(service service.AdminService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var book model.Book
		if err := c.BindJSON(&book); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "неверный формат данных книги"})
			return
		}

		bookID, err := service.AddBook(c.Request.Context(), book, c.GetString("user"))
		if err != nil {
			slog.Error(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "ошибка добавления книги"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"id": bookID})
	}
}
