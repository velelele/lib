package postgres

import (
	"context"
	"crud/internal/core/interface/repository"
	"crud/internal/core/model"
	"crud/internal/lib/db"
	"database/sql"
	"errors"
	"fmt"
)

type _favoriteRepository struct {
	db *db.Db
}

func NewFavoriteRepo(db *db.Db) repository.FavoriteRepository {
	return _favoriteRepository{db}
}

func (favoriteRepository _favoriteRepository) AddFavoriteBook(ctx context.Context, login string, bookID int) error {
	userID, err := favoriteRepository.GetId(ctx, login)

	_, err = favoriteRepository.db.PgConn.Exec(ctx, `INSERT INTO favorites_books (user_id, book_id) VALUES ($1, $2)`, userID, bookID)
	if err != nil {
		return fmt.Errorf("ошибка при добавлении книги в избранное: %s", err.Error())
	}
	return nil
}

func (favoriteRepository _favoriteRepository) GetFavoriteBooks(ctx context.Context, login string) ([]model.Book, error) {
	userID, err := favoriteRepository.GetId(ctx, login)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении user_id: %s", err.Error())
	}

	var books []model.Book

	rows, err := favoriteRepository.db.PgConn.Query(ctx, "SELECT b.title, b.author, b.body, b.release FROM books b JOIN favorites_books f ON b.id = f.book_id WHERE f.user_id = $1", userID)
	if err != nil {
		return nil, fmt.Errorf("ошибка получения списка избранных книг: %s", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var book model.Book
		if err := rows.Scan(&book.Title, &book.Author, &book.Body, &book.Release); err != nil {
			return nil, fmt.Errorf("ошибка сканирования строки: %s", err.Error())
		}
		books = append(books, book)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка при переборе строк: %s", err.Error())
	}

	return books, nil
}

func (favoriteRepository _favoriteRepository) GetId(ctx context.Context, login string) (int, error) {
	var userID int
	err := favoriteRepository.db.PgConn.QueryRow(ctx, "SELECT id FROM users WHERE login = $1", login).Scan(&userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, fmt.Errorf("пользователь с логином '%s' не найден", login)
		}
		return 0, fmt.Errorf("ошибка при поиске пользователя: %s", err.Error())
	}
	return userID, nil
}
