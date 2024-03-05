package http

import (
	"crud/internal/core/interface/service"
	"crud/internal/transport/handler"
	"crud/internal/transport/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoutes(service service.AuthService, postService service.PostService, bookService service.BookService) *gin.Engine {
	router := gin.Default()

	router.POST("/register", handler.RegisterUser(service))

	router.GET("/auth", handler.Auth(service))

	api := router.Group("/api", middleware.AuthMiddleware)
	{
		api.POST("/post", handler.CreatePost(postService))
		api.GET("/post/:id", handler.GetPost(postService))
		api.GET("/book", handler.SearchBooks(bookService))
		api.GET("/books", handler.GetAllBooks(bookService))

	}
	return router
}
