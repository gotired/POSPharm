package usecase

import (
	"github.com/gotired/POSPharm/internal/domain"
	"github.com/gotired/POSPharm/internal/repository"
)

type branchUsecase struct {
	repo repository.Branch
}

func NewBranch(repo repository.Branch) Branch {
	return &branchUsecase{repo: repo}
}

func (u *branchUsecase) List(limit int, offset int) ([]domain.ListBranch, error) {
	return u.repo.List(limit, offset)
}
