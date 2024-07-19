package controller

import (
	"errors"
	"lms-user/database"
	"lms-user/helper"
	"lms-user/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	PENDING  = "PENDING"
	APPROVED = "APPROVED"
	REJECTED = "REJECTED"
)

type LeaveRequest struct {
	Id            string `json:"id" db:"id"`
	ManagerId     string `json:"manager_id" db:"manager_id"`
	ApplicantId   string `json:"applicant_id" db:"applicant_id"`
	StartDatetime string `json:"start_datetime" db:"start_datetime"`
	EndDatetime   string `json:"end_datetime" db:"end_datetime"`
}

func CreateLeave(context *gin.Context) {
	var err error
	var request LeaveRequest

	if err := context.Bind(&request); err != nil {
		ErrorHandler(context, err)
		return
	}

	applicantId := helper.CurrentUser(context).ID

	managerId, err := strconv.ParseUint(request.ManagerId, 10, 32)
	if err != nil {
		ErrorHandler(context, err)
		return
	}

	const shortForm = "2006-Jan-02"
	starttime, err := time.Parse(shortForm, request.StartDatetime)
	if err != nil {
		ErrorHandler(context, err)
		return
	}
	endtime, err := time.Parse(shortForm, request.EndDatetime)
	if err != nil {
		ErrorHandler(context, err)
		return
	}

	leave := model.Leave{
		ApplicantID:   applicantId,
		ManagerID:     uint(managerId),
		Type:          "Holiday",
		Status:        PENDING,
		StartDatetime: starttime,
		EndDatetime:   endtime,
	}

	_, err = leave.Save()
	if err != nil {
		ErrorHandler(context, err)
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Success"})

}

func GetLeaves(context *gin.Context) {
	var request LeaveRequest
	var leaves []model.Leave
	if err := context.ShouldBindJSON(&request); err != nil {
		ErrorHandler(context, err)
		return
	}

	if request.ApplicantId == "" {
		ErrorHandler(context, errors.New("missing in request ApplicantId"))
		return
	}

	err := database.Database.
		Where("applicant_id=?", request.ApplicantId).
		Where("deleted_at IS NULL").
		Find(&leaves).Error

	if err != nil {
		ErrorHandler(context, err)
		return
	}

	context.JSON(http.StatusCreated, gin.H{"leaves": leaves})
}

func ApproveLeaveById(context *gin.Context) {
	var request LeaveRequest
	var leave model.Leave
	if err := context.Bind(&request); err != nil {
		ErrorHandler(context, err)
		return
	}

	currentUser := helper.CurrentUser(context)
	err := database.Database.
		Where("id=?", request.Id).
		Where("manager_id=?", currentUser.ID).
		Where("status=?", PENDING).
		First(&leave).Error
	if err != nil {
		ErrorHandler(context, err)
		return
	}

	leave.Status = APPROVED
	_, err = leave.Update()
	if err != nil {

		ErrorHandler(context, err)
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Success"})

}

func RejectLeaveById(context *gin.Context) {
	var request LeaveRequest
	var leave model.Leave
	if err := context.Bind(&request); err != nil {
		ErrorHandler(context, err)
		return
	}

	currentUser := helper.CurrentUser(context)
	if (currentUser == model.User{}) {
		ErrorHandler(context, errors.New("provide user"))
		return
	}

	managerId, err := strconv.ParseUint(request.ManagerId, 10, 32)
	if err != nil {
		ErrorHandler(context, err)
		return
	}
	if currentUser.ID != uint(managerId) {
		ErrorHandler(context, ErrNotFound)
		return
	}
	err = database.Database.
		Where("id=?", request.Id).
		Where("manager_id=?", request.ManagerId).
		Where("status=?", PENDING).
		First(&leave).Error
	if err != nil {
		ErrorHandler(context, err)
		return
	}

	leave.Status = REJECTED
	_, err = leave.Update()
	if err != nil {
		ErrorHandler(context, err)
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Success"})
}
