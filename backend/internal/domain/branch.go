package domain

import "gorm.io/gorm"

type Branch struct {
	gorm.Model
	Name string `json:"name"`
}

type ListBranch struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
