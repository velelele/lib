package repository

import (
	"context"
	"crud/internal/core/model"
)

type AuthRepository interface {
	GetUser(ctx context.Context, login, hashPassword string) (string, error)
	Register(ctx context.Context, login, hashPassword string) (string, error)
}

type PostRepository interface {
	CreatePost(ctx context.Context, post model.Post) (int, error)
	GetPost(ctx context.Context, postId int) (model.Post, error)
}

type BookRepository interface {
	GetBook(ctx context.Context, bookId int) (model.Book, error)
	SearchBooks(ctx context.Context, title, author, release string) ([]model.Book, error)
	GetAllBooks(ctx context.Context) ([]model.Book, error)
}
