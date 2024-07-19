package main

import (
	"fmt"
	"log"
	"net/http"

	"lms-user/controller"
	"lms-user/database"
	"lms-user/helper"
	"lms-user/model"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.User{}, &model.Leave{})

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

	//leaves
	leavesRoutes := router.Group("/leaves")
	leavesRoutes.Use(helper.JWTAuth())
	leavesRoutes.GET("", controller.GetLeaves)
	leavesRoutes.POST("/create", controller.CreateLeave)
	leavesRoutes.POST("/approve", controller.ApproveLeaveById)
	leavesRoutes.POST("/reject", controller.RejectLeaveById)

	//router.POST("/logout", )
	router.POST("/signup", controller.Signup)
	router.POST("/login", controller.Login)

	//manager
	leavesRoutes.GET("/view_user", controller.ViewUserToken)
	leavesRoutes.GET("/view_team", controller.ViewTeam)
	leavesRoutes.POST("/add_team", controller.AddUsertoTeam)
	leavesRoutes.GET("/pending_leaves", controller.ViewPendingTeamLeaves)

	router.GET("/ping", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "pong"}) })

	fmt.Println("Server running on port 8000")
	router.Run(":8000")
}
