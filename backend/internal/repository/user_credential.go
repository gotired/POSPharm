package repository

import (
	"github.com/gotired/POSPharm/internal/domain"
	"gorm.io/gorm"
)

type UserCredential interface {
	ValidateEmail(email string) (bool, error)
	ValidateUsername(username string) (bool, error)
	Register(tx *gorm.DB, userID uint, username string, email string, passwordHash string) error
	GetIDByUsernameEmail(username, email string) (*domain.UserLogin, error)
}
