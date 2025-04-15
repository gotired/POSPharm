package usecase

import "github.com/gotired/POSPharm/internal/domain"

type User interface {
	GetUsers() ([]domain.User, error)
	GetUser(id uint) (*domain.User, error)
	RegisterUser(detail domain.UserCreate, passwordHash string) error
	GetUserLogin(username string) (*domain.UserLogin, error)
	ValidateRegister(detail domain.UserCreate) error
}
