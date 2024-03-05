package postgres

import (
	"context"
	"crud/internal/core/interface/repository"
	"crud/internal/core/model"
	"crud/internal/lib/db"
	"crud/internal/repository/dbModel"
	"fmt"
)

type _bookRepository struct {
	db *db.Db
}

func NewBookRepo(db *db.Db) repository.BookRepository {
	return _bookRepository{db}
}

func (bookRepository _bookRepository) GetBook(ctx context.Context, bookId int) (model.Book, error) {
	var book dbModel.Book

	err := bookRepository.db.PgConn.QueryRow(ctx,
		`SELECT title, body, release, author FROM books WHERE id=$1`,
		bookId).Scan(&book.Title, &book.Body, &book.Release, &book.Author)

	if err != nil {
		return model.Book{}, fmt.Errorf("ошибка получения книги: %s", err.Error())
	}

	return model.Book(book), nil
}

func (bookRepository _bookRepository) SearchBooks(ctx context.Context, title, author, release string) ([]model.Book, error) {
	var books []dbModel.Book

	query := "SELECT title, body, release, author FROM books WHERE 1=1"

	if title != "" {
		query += " AND title ILIKE '%" + title + "%'"
	}
	if author != "" {
		query += " AND author ILIKE '%" + author + "%'"
	}
	if release != "" {
		query += " AND release ILIKE '%" + release + "%'"
	}

	rows, err := bookRepository.db.PgConn.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error retrieving books: %s", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var book dbModel.Book
		err := rows.Scan(&book.Title, &book.Body, &book.Release, &book.Author)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %s", err.Error())
		}
		books = append(books, book)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %s", err.Error())
	}

	var result []model.Book
	for _, b := range books {
		result = append(result, model.Book(b))
	}

	return result, nil
}

func (bookRepository _bookRepository) GetAllBooks(ctx context.Context) ([]model.Book, error) {
	var books []dbModel.Book

	rows, err := bookRepository.db.PgConn.Query(ctx, "SELECT title, body, release, author FROM books")
	if err != nil {
		return nil, fmt.Errorf("ошибка получения списка книг: %s", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var book dbModel.Book
		err := rows.Scan(&book.Title, &book.Body, &book.Release, &book.Author)
		if err != nil {
			return nil, fmt.Errorf("ошибка сканирования строки: %s", err.Error())
		}
		books = append(books, book)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка при переборе строк: %s", err.Error())
	}

	var result []model.Book
	for _, b := range books {
		result = append(result, model.Book(b))
	}

	return result, nil
}

func (bookRepository _bookRepository) AddBook(ctx context.Context, book model.Book) (int, error) {
	var bookID int
	err := bookRepository.db.PgConn.QueryRow(ctx,
		`INSERT INTO books (title, author, body, release) VALUES ($1, $2, $3, $4) RETURNING id`,
		book.Title, book.Author, book.Body, book.Release).Scan(&bookID)
	if err != nil {
		return 0, fmt.Errorf("ошибка при добавлении книги: %s", err.Error())
	}
	return bookID, nil
}
