package routes

import (
	"hometest/handlers"
	"hometest/packages/connection"
	"hometest/repositories"

	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.RouterGroup) {
	repo := repositories.MakeRepository(connection.DB)
	handler := handlers.HandlerUser(repo)

	r.POST("/user", handler.CreateUser)
}
