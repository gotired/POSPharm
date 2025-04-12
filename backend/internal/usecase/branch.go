package usecase

import "github.com/gotired/POSPharm/internal/domain"

type Branch interface {
	List(limit int, offset int) ([]domain.ListBranch, error)
}
