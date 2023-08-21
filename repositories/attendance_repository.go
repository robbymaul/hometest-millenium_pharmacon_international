package repositories

import "hometest/models"

type AttendanceRepository interface {
	ClockIn(attendance models.Attendance) (models.Attendance, error)
	ClockOut(attendance models.Attendance) (models.Attendance, error)
	GetAttendance(userId int) (models.Attendance, error)
}

func (r *repository) ClockIn(attendance models.Attendance) (models.Attendance, error) {
	var data models.Attendance
	err := r.db.Create(&attendance).Where("user_id =? ", attendance.UserID).Order("id DESC").Preload("User").First(&data).Error

	return data, err
}

func (r *repository) ClockOut(attendance models.Attendance) (models.Attendance, error) {
	err := r.db.Save(&attendance).Error

	return attendance, err
}

func (r *repository) GetAttendance(userId int) (models.Attendance, error) {
	var attendance models.Attendance
	err := r.db.Preload("User").First(&attendance, "user_id = ?", userId).Error

	return attendance, err
}
