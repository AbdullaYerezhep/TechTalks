package service

import (
	"forum/models"
	"forum/pkg/repository"
)

type PostService struct {
	repo repository.Post
}

func NewPost(repo repository.Post) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) CreatePost(p models.Post) error {
	return s.repo.CreatePost(p)
}

func (s *PostService) GetPost(id int) (models.Post, error) {
	return s.repo.GetPost(id)
}

func (s *PostService) GetAllPosts() ([]models.Post, error) {
	return s.repo.GetAllPosts()
}

func (s *PostService) UpdatePost(p models.Post) error {
	return s.repo.UpdatePost(p)
}

func (s *PostService) DeletePost(id int) error {
	return s.repo.DeletePost(id)
}
