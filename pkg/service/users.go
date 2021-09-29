package service

import (
	"api_learning/models"
	"api_learning/pkg/repository"
)

type UserObjectsService struct {
	repo repository.UserObjects
}

func NewUserObjectsService(repo repository.UserObjects) *UserObjectsService {
	return &UserObjectsService{repo: repo}
}


func (s *UserObjectsService) ListAll() ([]models.UserResponse, error) {
	users, err := s.repo.ListAll()
	if err != nil {
		return []models.UserResponse{}, err
	}
	return users, err
}