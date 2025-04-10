package repository

import "github.com/gotired/POSPharm/internal/domain"

type Role interface {
	List(limit int, offset int) ([]domain.ListRole, error)
}
