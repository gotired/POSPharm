package usecase

import "github.com/golang-jwt/jwt/v5"

type Auth interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword, password string) bool
	GenerateToken(userID uint) (string, error)
	GenerateRefreshToken(userID uint) (string, error)
	ValidateAccessToken(token string) (*jwt.MapClaims, error)
	ValidateRefreshToken(token string) (*jwt.MapClaims, error)
	EncryptToken(token string) (string, error)
	DecryptToken(token string) (string, error)
}
