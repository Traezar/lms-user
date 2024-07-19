package controller

import (
	"errors"
	"lms-user/database"
	"lms-user/helper"
	"lms-user/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AddtoTeamRequest struct {
	UserId string `json:"user_id" db:"user_id"`
}

func ViewTeam(context *gin.Context) {
	id := helper.CurrentUser(context).ID
	var team []model.User
	err := database.Database.Where("manager_id=?", id).Find(&team).Error
	if err != nil {
		ErrorHandler(context, err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"team": team})
}

func AddUsertoTeam(context *gin.Context) {
	manager := helper.CurrentUser(context)
	println("%d", manager.ID)

	var team []model.User
	var user model.User
	var request AddtoTeamRequest

	if err := context.Bind(&request); err != nil {
		ErrorHandler(context, err)
		return
	}
	if request.UserId == "" {
		ErrorHandler(context, errors.New("provide a user_id"))
		return
	}

	//from request
	userId, err := strconv.ParseUint(request.UserId, 10, 32)
	if err != nil {
		ErrorHandler(context, err)
		return
	}

	err = database.Database.Where("id=?", userId).Find(&user).Error
	if err != nil {
		ErrorHandler(context, err)
		return
	}

	user.ManagerID = manager.ID
	_, err = user.Update()
	if err != nil {
		ErrorHandler(context, err)
		return
	}

	err = database.Database.Where("manager_id=?", manager.ID).Find(&team).Error
	if err != nil {
		ErrorHandler(context, err)
		return
	}
	context.JSON(http.StatusOK, gin.H{"team": team})
}

func ViewPendingTeamLeaves(context *gin.Context) {
	id := helper.CurrentUser(context).ID
	var leaves []model.Leave
	err := database.Database.Where("manager_id=?", id).Where("status =?", "PENDING").Find(&leaves).Error
	if err != nil {
		ErrorHandler(context, err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"leaves": leaves})
}
