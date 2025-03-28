package repository

import "github.com/gotired/POSPharm/internal/domain"

type UserRepository interface {
	GetAllUsers() ([]domain.User, error)
	GetUserByID(id uint) (*domain.User, error)
	CreateUser(user *domain.User) error
}
