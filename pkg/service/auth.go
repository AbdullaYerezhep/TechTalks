package service

import (
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

func (s *AuthService) GetUser(name string) (models.User, error) {
	return s.repo.GetUser(name)
}

func (s *AuthService) GetUserByToken(token string) (models.User, error) {
	return s.repo.GetUserByToken(token)
}

func getHash(pass string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(pass), 0)
	return string(hash)
}
