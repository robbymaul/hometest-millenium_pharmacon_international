package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `json:"name" gorm:"type: varchar(255)"`
}

type UserResponse struct {
	gorm.Model
	Name string `json:"name"`
}

func (UserResponse) TableName() string {
	return "users"
}
