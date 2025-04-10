package domain

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name string `json:"name"`
}

type ListRole struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
