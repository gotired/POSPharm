package usecase

import (
	"errors"

	"github.com/gotired/POSPharm/internal/domain"
	"github.com/gotired/POSPharm/internal/repository"
	"gorm.io/gorm"
)

type userUsecase struct {
	db             *gorm.DB
	userRepo       repository.User
	userCredential repository.UserCredential
}

func NewUser(db *gorm.DB, userRepo repository.User, userCredentialRepo repository.UserCredential) User {
	return &userUsecase{db: db, userRepo: userRepo, userCredential: userCredentialRepo}
}

func (u *userUsecase) GetUsers() ([]domain.User, error) {
	return u.userRepo.List()
}

func (u *userUsecase) GetUser(id uint) (*domain.User, error) {
	return u.userRepo.GetByID(id)
}

func (u *userUsecase) RegisterUser(detail domain.UserCreate, passwordHash string) error {
	return u.db.Transaction(func(tx *gorm.DB) error {
		userID, err := u.userRepo.Register(tx, detail.FirstName, detail.LastName, detail.Tel)
		if err != nil {
			return err
		}
		if err = u.userCredential.Register(tx, *userID, detail.UserName, detail.Email, passwordHash); err != nil {
			return err
		}
		return nil
	})

}

func (u *userUsecase) GetUserLogin(username string) (*domain.UserLogin, error) {
	return u.userCredential.GetIDByUsernameEmail(username, username)
}

func (u *userUsecase) ValidateRegister(detail domain.UserCreate) error {
	if detail.Password != detail.ConfirmPassword {
		return errors.New("password is mismatch")
	}
	exists, err := u.userCredential.ValidateEmail(detail.Email)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("email is already exists")
	}
	exists, err = u.userCredential.ValidateUsername(detail.UserName)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("username is already exists")
	}
	exists, err = u.userRepo.ValidateTel(detail.Tel)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("tel is already exists")
	}
	return nil
}
