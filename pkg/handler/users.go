package handler

import (
	"api_learning/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) ListAll(c *gin.Context) {


	users, err := h.services.UserObjects.ListAll()

	if err!= nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllUsersResponse{
		Data: users,
	})
}


type getAllUsersResponse struct{
	Data []models.UserResponse "json: 'data'"
}