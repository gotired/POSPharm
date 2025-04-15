package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Tel       string `json:"tel" gorm:"unique"`
}

type UserCredential struct {
	UserID       uint   `gorm:"not null"`
	Username     string `json:"username" gorm:"unique"`
	Email        string `json:"email" gorm:"unique"`
	PasswordHash string `json:"-"`

	User User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type UserLogin struct {
	UserID       uint
	Email        string
	PasswordHash string
}

type UserBranchRole struct {
	UserID   uint `gorm:"not null"`
	BranchID uint `gorm:"not null"`
	RoleID   uint `gorm:"not null"`

	User   User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Branch Branch `gorm:"foreignKey:BranchID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Role   Role   `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type UserCreate struct {
	UserName        string `json:"username"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Email           string `json:"email"`
	Tel             string `json:"tel"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}
