package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name         string `json:"name"`
	Email        string `json:"email" gorm:"unique"`
	PasswordHash string `json:"-"`
	RoleID       uint   `json:"role_id"`
	Role         Role   `json:"role" gorm:"foreignKey:RoleID;references:ID"`
}

type UserLogin struct {
	Email        string
	PasswordHash string
}
