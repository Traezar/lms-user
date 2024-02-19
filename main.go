package main

import (
	"fmt"
	"log"
	"net/http"

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
	router.Static("/views", "./public/views")
	router.LoadHTMLGlob("./public/views/*html")
	// Routes
	router.GET("/signup", getSignupForm)
	router.GET("/login", getLoginForm)

	router.POST("/logout", getLoginForm)
	router.POST("/signup", controller.SignupForm)
	router.POST("/login", controller.LoginForm)

	router.POST("/register", controller.Signup)
	router.GET("/ping", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "pong"}) })

	fmt.Println("Server running on port 8000")
	router.Run(":8000")
}

func getSignupForm(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "signup.html", nil)
}
func getLoginForm(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", nil)
}
