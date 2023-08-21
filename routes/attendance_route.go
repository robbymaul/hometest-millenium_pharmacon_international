package routes

import (
	"hometest/handlers"
	"hometest/packages/connection"
	"hometest/packages/middleware"
	"hometest/repositories"

	"github.com/gin-gonic/gin"
)

func AttendanceRoute(r *gin.RouterGroup) {
	repo := repositories.MakeRepository(connection.DB)
	handler := handlers.HandlerAttendance(repo)

	r.POST("/clock-in", middleware.ClockInImage(), handler.ClockIn)
	r.PATCH("clock-out", middleware.ClockOutImage(), handler.ClockOut)
}
