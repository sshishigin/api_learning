package repository

import (
	"api_learning/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type UserObjects interface {
	ListAll() ([]models.UserResponse, error)
}

type Repository struct {
	Authorization
	UserObjects
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		UserObjects: NewUserObjectsPostgres(db),
	}
}
