package repository

import (
	"github.com/gotired/POSPharm/internal/domain"
	"gorm.io/gorm"
)

type userRepoImpl struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) User {
	return &userRepoImpl{db: db}
}

func (r *userRepoImpl) List() ([]domain.User, error) {
	var users []domain.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepoImpl) GetByID(id uint) (*domain.User, error) {
	var user domain.User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *userRepoImpl) Register(tx *gorm.DB, firstName, lastName, tel string) (*uint, error) {
	user := domain.User{
		FirstName: firstName,
		LastName:  lastName,
		Tel:       tel,
	}
	result := tx.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user.ID, nil
}

func (r *userRepoImpl) ValidateTel(tel string) (bool, error) {
	var count int64
	err := r.db.Model(&domain.User{}).Where("tel = ?", tel).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, err
}
