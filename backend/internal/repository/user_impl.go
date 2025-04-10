package repository

import (
	"github.com/gotired/POSPharm/internal/domain"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) User {
	return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) List() ([]domain.User, error) {
	var users []domain.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepositoryImpl) GetByID(id uint) (*domain.User, error) {
	var user domain.User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *userRepositoryImpl) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *userRepositoryImpl) GetByUsername(username string) (*domain.UserLogin, error) {
	var user domain.UserLogin
	err := r.db.Model(&domain.User{}).Select("email", "password_hash").First(&user).Error
	return &user, err
}
