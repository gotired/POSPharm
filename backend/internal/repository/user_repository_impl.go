package repository

import (
	"github.com/gotired/POSPharm/internal/domain"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) GetAllUsers() ([]domain.User, error) {
	var users []domain.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepositoryImpl) GetUserByID(id uint) (*domain.User, error) {
	var user domain.User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *userRepositoryImpl) CreateUser(user *domain.User) error {
	return r.db.Create(user).Error
}
