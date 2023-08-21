package main

import (
	"hometest/config"
	"hometest/packages/connection"
	"hometest/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	r := gin.Default()

	connection.Database()

	config.Migration()

	routes.RouteInit(r.Group("hometest"))

	r.Run("localhost:5000")
}
