package repository

import (
	"github.com/gotired/POSPharm/internal/domain"
	"gorm.io/gorm"
)

type branchRepositoryImpl struct {
	db *gorm.DB
}

func NewBranch(db *gorm.DB) Branch {
	return &branchRepositoryImpl{db: db}
}

func (r *branchRepositoryImpl) List(limit int, offset int) ([]domain.ListBranch, error) {
	var branches []domain.ListBranch
	err := r.db.Model(&domain.Branch{}).Select("id", "name").Offset(offset).Limit(limit).Find(&branches).Error
	return branches, err
}
