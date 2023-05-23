package service

import (
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
func (s *AuthService) CreateUser(u models.User) error {
	if err := validateNewUserData(u); err != nil {
		return err
	}
	u.Password = getHash(u.Password)
	if err := s.repo.CreateUser(u); err != nil {
		if strings.Contains(err.Error(), "UNIQUE") {
			return ErrUserExists
		} else {
			return err
		}
	}

	return nil
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
