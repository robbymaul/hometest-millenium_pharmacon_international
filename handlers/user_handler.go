package handlers

import (
	"hometest/dto"
	"hometest/models"
	"hometest/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userHandler struct {
	UserRepository repositories.UserRepository
}

func HandlerUser(UserRepository repositories.UserRepository) *userHandler {
	return &userHandler{UserRepository}
}

func (h *userHandler) CreateUser(c *gin.Context) {
	request := new(dto.UserRequestDTO)
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "ivalid input",
		})
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "ivalid input",
		})
		return
	}

	user := models.User{
		Name: request.Name,
	}

	err = h.UserRepository.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success create user",
	})
}
