package usecase

import "github.com/gotired/POSPharm/internal/domain"

type User interface {
	GetUsers() ([]domain.User, error)
	GetUser(id uint) (*domain.User, error)
	RegisterUser(user *domain.User) error
	GetUserLogin(username string) (*domain.UserLogin, error)
}
