package usecase

import (
	"github.com/gotired/POSPharm/internal/domain"
	"github.com/gotired/POSPharm/internal/repository"
)

type userUsecase struct {
	userRepo repository.User
}

func NewUser(repo repository.User) User {
	return &userUsecase{userRepo: repo}
}

func (u *userUsecase) GetUsers() ([]domain.User, error) {
	return u.userRepo.List()
}

func (u *userUsecase) GetUser(id uint) (*domain.User, error) {
	return u.userRepo.GetByID(id)
}

func (u *userUsecase) RegisterUser(user *domain.User) error {
	return u.userRepo.Create(user)
}

func (u *userUsecase) GetUserLogin(username string) (*domain.UserLogin, error) {
	return u.userRepo.GetByUsername(username)
}
