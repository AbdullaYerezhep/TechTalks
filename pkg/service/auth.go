package service

import (
	"fmt"
	"forum/models"
	"forum/pkg/repository"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuth(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

// New user actions
func (s *AuthService) CreateUser(u models.User) (int, error) {
	if err := validateNewUserData(u); err != nil {
		return 0, fmt.Errorf("failed to insert post: %v", err)
	}
	u.Password = getHash(u.Password)

	id, err := s.repo.CreateUser(u)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE") {
			return 0, fmt.Errorf("Error creating user: %v", ErrUserExists)
		} else {
			return 0, fmt.Errorf("Error creating user: %v", err)
		}
	}

	return id, nil
}

func (s *AuthService) GetUserByID(id int) (models.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *AuthService) GetUser(name, password string) (models.User, error) {
	user, err := s.repo.GetUser(name)
	if err != nil {
		return models.User{}, ErrUserNotFound
	} else if !verifyPass(password, user.Password) {
		return models.User{}, ErrWrongPassword
	}
	return user, nil
}

func getHash(pass string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(pass), 0)
	return string(hash)
}

func verifyPass(pass, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
}
