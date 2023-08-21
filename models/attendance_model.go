package models

import (
	"time"

	"gorm.io/gorm"
)

type Attendance struct {
	gorm.Model
	UserID        int       `json:"userId" gorm:"type: int"`
	User          User      `json:"user"`
	ClockIn       time.Time `json:"clockIn" gorm:"type: datetime"`
	ClockInImage  string    `json:"clockInImage" gorm:"type: varchar(255)"`
	ClockOut      time.Time `json:"clockOut" gorm:"type: datetime"`
	ClockOutImage string    `json:"clockOutImage" gorm:"type: varchar(255)"`
}
