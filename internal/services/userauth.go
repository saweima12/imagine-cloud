package services

import (
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"errors"

	"github.com/saweima12/imagine/internal/modules/config"
)

type UserAuthService interface {
	VerifyUser(username, password string) (bool, error)
	CheckAuthorization(username, password string) (bool, error)
	GenerateToken(username, password string) string
}

type service struct {
}

func NewUserAuthService() UserAuthService {
	return &service{}
}

func (s *service) VerifyUser(username, password string) (bool, error) {
	userContext := config.GetUserContext()

	if subtle.ConstantTimeCompare([]byte(username), []byte(userContext.Username)) == 1 &&
		subtle.ConstantTimeCompare([]byte(password), []byte(userContext.Password)) == 1 {
		return true, nil
	}
	return false, nil
}

func (s *service) CheckAuthorization(username, password string) (bool, error) {
	userContext := config.GetUserContext()

	hashString := s.GenerateToken(userContext.Username, userContext.Password)

	if subtle.ConstantTimeCompare([]byte(username), []byte(userContext.Username)) == 1 &&
		subtle.ConstantTimeCompare([]byte(password), []byte(hashString)) == 1 {
		return true, nil
	}

	err := errors.New("password not match")

	return false, err
}

func (s *service) GenerateToken(username, password string) string {
	// generate sha256 token return.
	hash := sha256.Sum256([]byte(username + password))
	hashString := hex.EncodeToString(hash[:])

	return hashString
}
