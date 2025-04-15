package repository

import (
	"fmt"

	"github.com/gotired/POSPharm/internal/domain"
	"gorm.io/gorm"
)

type userCredentialRepoImpl struct {
	db *gorm.DB
}

func NewUserCredential(db *gorm.DB) UserCredential {
	return &userCredentialRepoImpl{db: db}
}

func (r *userCredentialRepoImpl) validate(key string, value any) (bool, error) {
	var count int64
	err := r.db.Model(&domain.UserCredential{}).Where(fmt.Sprintf("%s = ?", key), value).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, err
}

func (r *userCredentialRepoImpl) ValidateEmail(email string) (bool, error) {
	return r.validate("email", email)
}

func (r *userCredentialRepoImpl) ValidateUsername(username string) (bool, error) {
	return r.validate("username", username)
}

func (r *userCredentialRepoImpl) Register(tx *gorm.DB, userID uint, username string, email string, passwordHash string) error {
	userCred := domain.UserCredential{
		UserID:       userID,
		Email:        email,
		Username:     username,
		PasswordHash: passwordHash,
	}
	return tx.Create(&userCred).Error
}

func (r *userCredentialRepoImpl) GetIDByUsernameEmail(username, email string) (*domain.UserLogin, error) {
	var user domain.UserLogin
	result := r.db.Model(&domain.UserCredential{}).Where("username = ?", username).Or("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
