package service

import (
	"forum/models"
	"forum/pkg/repository"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuth(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(u *models.User) error {
	u.Password = getHash(u.Password)
	return s.repo.CreateUser(u)
}

func (s *AuthService) GetByName(name string) (*models.User, error) {
	return s.repo.GetByName(name)
}
