package handlers

import (
	"hometest/dto"
	"hometest/models"
	"hometest/repositories"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	uploadsDir          = "uploads"
	similarityThreshold = 0.8
)

type attendanceHandler struct {
	AttendanceRepositor repositories.AttendanceRepository
}

func HandlerAttendance(AttendanceRepository repositories.AttendanceRepository) *attendanceHandler {
	return &attendanceHandler{AttendanceRepository}
}

func (h *attendanceHandler) ClockIn(c *gin.Context) {
	var err error

	userId, _ := strconv.Atoi(c.PostForm("userId"))
	imageFile, _ := c.Get("clockInImage")
	tempImagePath, _ := c.Get("clockInImage")

	similarity := CompareImage(imageFile.(string), tempImagePath.(string))
	if !similarity {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Attendance submission is invalid. Please retake the photo.",
		})
		return
	}

	request := dto.ClockInRequstDTO{
		UserID:       userId,
		ClockInImage: imageFile.(string),
	}

	attendance := models.Attendance{
		UserID:        request.UserID,
		ClockIn:       time.Now(),
		ClockInImage:  request.ClockInImage,
		ClockOut:      time.Time{},
		ClockOutImage: "",
	}

	data, err := h.AttendanceRepositor.ClockIn(attendance)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   data,
	})
}

func (h *attendanceHandler) ClockOut(c *gin.Context) {
	userId, _ := strconv.Atoi(c.PostForm("userId"))
	imageFile, _ := c.Get("clockOutImage")
	tempImagePath, _ := c.Get("clockOutImage")

	similarity := CompareImage(imageFile.(string), tempImagePath.(string))
	if !similarity {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Attendance submission is invalid. Please retake the photo.",
		})
		return
	}

	request := dto.ClockOutRequestDTO{
		ClockOutImage: imageFile.(string),
	}

	attendance, _ := h.AttendanceRepositor.GetAttendance(userId)

	if request.ClockOutImage != "" {
		attendance.ClockOutImage = request.ClockOutImage
	}

	attendance.ClockOut = time.Now()

	data, _ := h.AttendanceRepositor.ClockOut(attendance)

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   data,
	})

}

func CompareImage(imageFile, tempImagePath string) bool {
	similarity := calculateStringSimilarity(imageFile, tempImagePath)
	return similarity >= similarityThreshold
}

func calculateStringSimilarity(str1, str2 string) float64 {
	if str1 == str2 {
		return 1.0
	}
	return 0.0
}
