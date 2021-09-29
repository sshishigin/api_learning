package service

import (
	"api_learning/models"
	"api_learning/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type UserObjects interface {
	ListAll() ([]models.UserResponse, error)
	//Detail() (models.User, error)
}


type Service struct {
	Authorization
	UserObjects
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		UserObjects : NewUserObjectsService(repos.UserObjects),
	}
}
