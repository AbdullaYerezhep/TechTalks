package service

import (
	"errors"
	"forum/models"
	"forum/pkg/repository"
)

type PostService struct {
	repo    repository.Post
	repocat repository.Category
}

func NewPost(repo repository.Post, repocat repository.Category) *PostService {
	return &PostService{
		repo:    repo,
		repocat: repocat,
	}
}

func (s *PostService) CreatePost(p models.Post) error {
	cats, err := s.repocat.GetCategories()
	if err != nil {
		return err
	}
	if !containsAll(p.Category, cats) {
		return errors.New("invalid category")
	}
	return s.repo.CreatePost(p)
}

func (s *PostService) GetCategories() ([]string, error) {
	return s.repocat.GetCategories()
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

func (s *PostService) DeletePost(user_id, post_id int) error {
	return s.repo.DeletePost(user_id, post_id)
}

func (s *PostService) RatePost(rate models.RatePost) error {
	return s.repo.LikeDis(rate)
}

func containsAll(list, target []string) bool {
	for i := range list {
		contains := false
		for j := range target {
			if list[i] == target[j] {
				contains = true
				break
			}
		}
		if !contains {
			return false
		}
	}
	return true
}
