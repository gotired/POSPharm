package usecase

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type authRepositoryImpl struct {
	accessSecretKey  []byte
	refreshSecretKey []byte
	aesKey           []byte
}

func NewAuth() Auth {
	return &authRepositoryImpl{accessSecretKey: []byte(os.Getenv("ACCESS_TOKEN")), refreshSecretKey: []byte(os.Getenv("REFRESH_TOKEN")), aesKey: []byte(os.Getenv("AES_KEY"))}
}

func (r *authRepositoryImpl) generate(userID uint, exp int64, key []byte) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     exp,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}

func (r *authRepositoryImpl) validate(token string, key []byte) (*jwt.MapClaims, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return key, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil || !parsedToken.Valid {
		return nil, errors.New("invalid token")
	}

	tokenDetail, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	return &tokenDetail, nil
}

func (r *authRepositoryImpl) HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedBytes), err
}

func (r *authRepositoryImpl) ComparePassword(hashedPassword, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}

func (r *authRepositoryImpl) GenerateToken(userID uint) (string, error) {
	return r.generate(userID, time.Now().Add(time.Hour*1).Unix(), r.accessSecretKey)
}

func (r *authRepositoryImpl) GenerateRefreshToken(userID uint) (string, error) {
	return r.generate(userID, time.Now().Add(time.Hour*24).Unix(), r.refreshSecretKey)
}

func (r *authRepositoryImpl) ValidateAccessToken(token string) (*jwt.MapClaims, error) {
	return r.validate(token, r.accessSecretKey)
}

func (r *authRepositoryImpl) ValidateRefreshToken(token string) (*jwt.MapClaims, error) {
	return r.validate(token, r.refreshSecretKey)
}

func (r *authRepositoryImpl) hashKey(key []byte) []byte {
	hash := sha256.Sum256(key)
	return hash[:]
}

func (r *authRepositoryImpl) EncryptToken(token string) (string, error) {
	block, err := aes.NewCipher(r.hashKey(r.aesKey))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	cipherText := gcm.Seal(nonce, nonce, []byte(token), nil)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func (r *authRepositoryImpl) DecryptToken(token string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(r.hashKey(r.aesKey))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	if len(data) < gcm.NonceSize() {
		return "", errors.New("ciphertext too short")
	}

	nonce, cipherTextBytes := data[:gcm.NonceSize()], data[gcm.NonceSize():]
	plainText, err := gcm.Open(nil, nonce, cipherTextBytes, nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}
