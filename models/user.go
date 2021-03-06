package models
import (
	_ "encoding/json"
)
type User struct {
	ID       int    `json:"id" `
	Name     string `json:"name" binding:required`
	Username string `json:"username" binding:required`
	Password string `json:"password" binding:required`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name" binding:required`
	Username string `json:"username" binding:required`
}
