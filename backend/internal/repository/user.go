package repository

import "github.com/gotired/POSPharm/internal/domain"

type User interface {
	List() ([]domain.User, error)
	GetByID(id uint) (*domain.User, error)
	Create(user *domain.User) error
	GetByUsername(username string) (*domain.UserLogin, error)
}
