package service

import (
	"forum/models"
	"forum/pkg/repository"
)

type SessionService struct {
	repo repository.Session
}

func NewSession(repo repository.Session) *SessionService {
	return &SessionService{
		repo: repo,
	}
}

func (s *SessionService) CreateSession(session models.Session) error {
	return s.repo.CreateSession(session)
}

func (s *SessionService) GetSession(token string) (models.Session, error) {
	return s.repo.GetSession(token)
}

func (s *SessionService) DeleteSession(user_id int) error {
	return s.repo.DeleteSession(user_id)
}
