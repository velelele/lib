package postgres

import (
	"context"
	"crud/internal/core/interface/repository"
	"crud/internal/core/model"
	"crud/internal/lib/db"
	"fmt"
)

type _adminRepository struct {
	db *db.Db
}

func NewAdminRepo(db *db.Db) repository.AdminRepository {
	return _adminRepository{db}
}

func (adminRepository _adminRepository) AddBook(ctx context.Context, book model.Book) (int, error) {
	var bookID int
	err := adminRepository.db.PgConn.QueryRow(ctx,
		`INSERT INTO books (title, author, body, release) VALUES ($1, $2, $3, $4) RETURNING id`,
		book.Title, book.Author, book.Body, book.Release).Scan(&bookID)
	if err != nil {
		return 0, fmt.Errorf("ошибка при добавлении книги: %s", err.Error())
	}
	return bookID, nil
}

func (adminRepository _adminRepository) GetAdmin(ctx context.Context, login string) (bool, error) {

	var admin bool
	err := adminRepository.db.PgConn.QueryRow(ctx,
		`SELECT admin from users where login = $1`, login).Scan(&admin)
	if err != nil {
		return false, fmt.Errorf("ошибка при получения поля admin: %s", err.Error())
	}
	return admin, nil
}
