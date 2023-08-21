package config

import (
	"fmt"
	"hometest/models"
	"hometest/packages/connection"
)

func Migration() {
	err := connection.DB.AutoMigrate(
		&models.User{},
		&models.Attendance{},
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Migration success")
}
