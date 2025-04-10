package usecase

type Auth interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword, password string) bool
}
