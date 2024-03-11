package handler

import (
	"crud/internal/core/interface/service"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

func AddFavoriteBook(service service.FavoriteService) gin.HandlerFunc {
	return func(c *gin.Context) {
		login := c.GetString("user")

		bookID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "неверный формат ID книги"})
			return
		}

		err = service.AddFavoriteBook(c.Request.Context(), login, bookID)
		if err != nil {
			slog.Error(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "ошибка при добавлении книги в избранное"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Книга успешно добавлена в избранное"})
	}
}

func GetFavoriteBooks(service service.FavoriteService) gin.HandlerFunc {
	return func(c *gin.Context) {
		login := c.GetString("user")

		books, err := service.GetFavoriteBooks(c.Request.Context(), login)
		if err != nil {
			slog.Error(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "ошибка при получении списка избранных книг"})
			return
		}

		c.JSON(http.StatusOK, books)
	}
}
