package repository

import (
	"github.com/gotired/POSPharm/internal/domain"
	"gorm.io/gorm"
)

type User interface {
	List() ([]domain.User, error)
	GetByID(id uint) (*domain.User, error)
	Register(tx *gorm.DB, firstName, lastName, tel string) (*uint, error)
	ValidateTel(tel string) (bool, error)
}
