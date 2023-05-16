package service

import (
	"forum/models"
	"forum/pkg/repository"
)

type CommentService struct {
	repo repository.Comment
}

func NewComment(repo repository.Comment) *CommentService {
	return &CommentService{repo: repo}
}

func (s *CommentService) AddComment(com models.Comment) error {
	return s.repo.AddComment(com)
}

func (s *CommentService) GetComment(id int) (models.Comment, error) {
	return s.repo.GetComment(id)
}

func (s *CommentService) GetPostComments(id int) ([]models.Comment, error) {
	return s.repo.GetPostComments(id)
}

func (s *CommentService) UpdateComment(com models.Comment) error {
	return s.repo.UpdateComment(com)
}

func (s *CommentService) DeleteComment(id int) error {
	return s.repo.DeleteComment(id)
}

func (s *CommentService) RateComment(rate models.RateComment) error {
	return s.repo.RateComment(rate)
}
