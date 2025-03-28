package usecase

import "github.com/gotired/POSPharm/internal/domain"

type UserUsecase interface {
	GetUsers() ([]domain.User, error)
	GetUser(id uint) (*domain.User, error)
	RegisterUser(user *domain.User) error
}
