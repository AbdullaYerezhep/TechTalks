package service

import (
	"forum/models"
	"forum/pkg/repository"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type Authorization interface {
	CreateUser(user *models.User) error
	GetByName(name string) (*models.User, error)
}

type Service struct {
	Authorization
}

func New(repo *repository.Repository) *Service {
	return &Service{
		Authorization: repo,
	}
}

func getHash(pass string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), 0)
	if err != nil {
		log.Fatal("Hashing error")
	}
	return string(hash)
}
