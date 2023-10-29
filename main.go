package main

import (
	"fmt"
	"log"

	"lms-user/controller"
	"lms-user/database"
	"lms-user/model"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.User{})

}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	loadEnv()
	loadDatabase()
	runApplication()
}

func runApplication() {
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Signup)
	publicRoutes.POST("/login", controller.Login)
	router.Run(":8000")
	fmt.Println("Server running on port 8000")

}
