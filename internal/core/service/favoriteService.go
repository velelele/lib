package service

import (
	"context"
	"crud/internal/core/interface/repository"
	"crud/internal/core/interface/service"
	"crud/internal/core/model"
)

type _favoriteService struct {
	repo repository.FavoriteRepository
}

func NewFavoriteService(repo repository.FavoriteRepository) service.FavoriteService {
	return _favoriteService{repo: repo}
}

func (favoriteService _favoriteService) AddFavoriteBook(ctx context.Context, login string, bookID int) error {
	return favoriteService.repo.AddFavoriteBook(ctx, login, bookID)
}

func (favoriteService _favoriteService) GetFavoriteBooks(ctx context.Context, login string) ([]model.Book, error) {
	return favoriteService.repo.GetFavoriteBooks(ctx, login)
}
