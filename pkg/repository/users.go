package repository

import (
	"api_learning/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UserObjectsPostgres struct {
	db *sqlx.DB
}

func NewUserObjectsPostgres(db *sqlx.DB) *UserObjectsPostgres {
	return &UserObjectsPostgres{db: db}
}


func (r *UserObjectsPostgres) ListAll () (users []models.UserResponse, err error) {
	query := fmt.Sprintf("SELECT id, name, username FROM %s \n", usersTable)
	err = r.db.Select(&users, query)
	return
}