package usecase

import (
	"github.com/gotired/POSPharm/internal/domain"
	"github.com/gotired/POSPharm/internal/repository"
)

type roleUsecase struct {
	roleRepo repository.Role
}

func NewRole(repo repository.Role) Role {
	return &roleUsecase{roleRepo: repo}
}

func (u *roleUsecase) List(limit int, offset int) ([]domain.ListRole, error) {
	return u.roleRepo.List(limit, offset)
}
