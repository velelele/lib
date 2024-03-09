package service

import (
	"context"
	"crud/internal/core/interface/repository"
	"crud/internal/core/interface/service"
	"crud/internal/core/model"
	"errors"
	"fmt"
	"log/slog"
)

type _adminService struct {
	repo repository.AdminRepository
}

func NewAdminService(repo repository.AdminRepository) service.AdminService {
	return _adminService{repo: repo}
}

func (adminService _adminService) AddBook(ctx context.Context, book model.Book, login string) (int, error) {
	admin, err := adminService.repo.GetAdmin(ctx, login)
	if err != nil {
		slog.Error(err.Error())
		return 0, errors.New("не смогли получить админа")
	}
	if admin {
		return adminService.repo.AddBook(ctx, book)
	}
	return 0, fmt.Errorf("доступ запрещён")
}
