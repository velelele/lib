package service

import (
	"context"
	"crud/internal/core/interface/repository"
	"crud/internal/core/interface/service"
	"crud/internal/core/model"
)

type _bookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) service.BookService {
	return _bookService{repo: repo}
}

func (bookService _bookService) GetBook(ctx context.Context, bookId int) (model.Book, error) {
	return bookService.repo.GetBook(ctx, bookId)
}

func (bookService _bookService) SearchBooks(ctx context.Context, title, author, release string) ([]model.Book, error) {
	return bookService.repo.SearchBooks(ctx, title, author, release)
}

func (bookService _bookService) GetAllBooks(ctx context.Context) ([]model.Book, error) {
	return bookService.repo.GetAllBooks(ctx)
}

func (bookService _bookService) AddBook(ctx context.Context, book model.Book) (int, error) {
	return bookService.repo.AddBook(ctx, book)
}
