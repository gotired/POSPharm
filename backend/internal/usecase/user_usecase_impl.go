package usecase

import (
	"github.com/gotired/POSPharm/internal/domain"
	"github.com/gotired/POSPharm/internal/repository"
)

type userUsecaseImpl struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecaseImpl{userRepo: repo}
}

func (u *userUsecaseImpl) GetUsers() ([]domain.User, error) {
	return u.userRepo.GetAllUsers()
}

func (u *userUsecaseImpl) GetUser(id uint) (*domain.User, error) {
	return u.userRepo.GetUserByID(id)
}

func (u *userUsecaseImpl) RegisterUser(user *domain.User) error {
	return u.userRepo.CreateUser(user)
}
