package repository

import (
	"crud/internal/core/interface/repository"
	"crud/internal/lib/db"
	"crud/internal/repository/postgres"
)

type RepositoryManager struct {
	repository.AuthRepository
	repository.PostRepository
	repository.BookRepository
	repository.AdminRepository
}

func NewRepositoryManager(db *db.Db) RepositoryManager {
	return RepositoryManager{
		postgres.NewRepo(db),
		postgres.NewPostRepo(db),
		postgres.NewBookRepo(db),
		postgres.NewAdminRepo(db),
	}
}
