package service

import (
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"errors"

	"github.com/saweima12/imagine/internal/imagine/config"
)

type UserAuthService interface {
	VerifyUser(username, password string) (string, error)
	CheckAuthorization(username, password string) bool
	GenerateToken(username, password string) string
}

type userAuthService struct {
	UserContext config.UserContext
}

func NewUserAuthService(userContext config.UserContext) UserAuthService {
	return &userAuthService{
		UserContext: userContext,
	}
}

func (s *userAuthService) VerifyUser(username, password string) (string, error) {
	userContext := s.UserContext

	if subtle.ConstantTimeCompare([]byte(username), []byte(userContext.Username)) == 1 &&
		subtle.ConstantTimeCompare([]byte(password), []byte(userContext.Password)) == 1 {
		// Verification successful, Generate Token password.
		hashString := s.GenerateToken(userContext.Username, userContext.Password)
		return hashString, nil
	}

	err := errors.New("password not match")

	return "", err
}

func (s *userAuthService) CheckAuthorization(username, password string) bool {
	userContext := s.UserContext

	hashString := s.GenerateToken(userContext.Username, userContext.Password)
	// compare user input & config.
	if subtle.ConstantTimeCompare([]byte(username), []byte(userContext.Username)) == 1 &&
		subtle.ConstantTimeCompare([]byte(password), []byte(hashString)) == 1 {
		return true
	}
	return false
}

func (s *userAuthService) GenerateToken(username, password string) string {
	// generate sha256 token return.
	hash := sha256.Sum256([]byte(username + password))
	hashString := hex.EncodeToString(hash[:])

	return hashString
}
