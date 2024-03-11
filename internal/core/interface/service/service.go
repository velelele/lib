package service

import (
	"context"
	"crud/internal/core/model"
)

type AuthService interface {
	Register(ctx context.Context, login, password string) (string, error)
	GenerateToken(ctx context.Context, login, password string) (string, error)
	Auth(ctx context.Context, login, password string) (string, error)
}

type PostService interface {
	CreatePost(ctx context.Context, post model.Post) (int, error)
	GetPost(ctx context.Context, postId int) (model.Post, error)
}

type BookService interface {
	GetBook(ctx context.Context, bookId int) (model.Book, error)
	SearchBooks(ctx context.Context, title, author, release string) ([]model.Book, error)
	GetAllBooks(ctx context.Context) ([]model.Book, error)
}

type AdminService interface {
	AddBook(ctx context.Context, book model.Book, login string) (int, error)
}

type FavoriteService interface {
	AddFavoriteBook(ctx context.Context, login string, bookID int) error
	GetFavoriteBooks(ctx context.Context, login string) ([]model.Book, error)
}
