package repository

import (
	"github.com/gotired/POSPharm/internal/domain"
	"gorm.io/gorm"
)

type roleRepositoryImpl struct {
	db *gorm.DB
}

func NewRole(db *gorm.DB) Role {
	return &roleRepositoryImpl{db: db}
}

func (r *roleRepositoryImpl) List(limit int, offset int) ([]domain.ListRole, error) {
	var roles []domain.ListRole
	err := r.db.Model(&domain.Role{}).Select("id", "name").Offset(offset).Limit(limit).Find(&roles).Error
	return roles, err
}
