package service

import (
	"errors"
	"forum/models"
	"forum/pkg/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuth(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(u models.User) error {
	u.Password = getHash(u.Password)
	return s.repo.CreateUser(u)
}

func (s *AuthService) GetUserByID(id int) (models.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *AuthService) GetUser(name, password string) (models.User, error) {
	user, err := s.repo.GetUser(name)
	if err != nil || !verifyPass(password, user.Password) {
		return user, errors.New("user not found")
	}
	return user, nil
}

func (s *AuthService) GetUserByToken(token string) (models.User, error) {
	return s.repo.GetUserByToken(token)
}

func getHash(pass string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(pass), 0)
	return string(hash)
}

func verifyPass(pass, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
}
