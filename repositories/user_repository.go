package repositories

import "hometest/models"

type UserRepository interface {
	CreateUser(user models.User) error
}

func (r *repository) CreateUser(user models.User) error {
	err := r.db.Create(&user).Error

	return err
}
