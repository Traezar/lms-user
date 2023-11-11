package controller

import (
	"lms-user/helper"
	"lms-user/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignupForm(context *gin.Context) {
	var input model.SignupInput

	if err := context.Bind(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := model.User{
		Name:        input.Name,
		Password:    input.Password,
		Email:       input.Email,
		Phonenumber: input.Phonenumber,
		Country:     input.Country,
		Gender:      input.Gender,
	}

	_, err := user.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Success"})
}

func LoginForm(context *gin.Context) {
	var input model.LoginInput
	println("login form")
	if err := context.Bind(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := model.FindUserByName(input.Name)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.ValidatePassword(input.Password)
	println("Got User: %s, from %s", user.Name, user.Country)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.HTML(http.StatusOK, "thanks.html", gin.H{
		"name":    user.Name,
		"country": user.Country,
	})
}

func Signup(context *gin.Context) {
	var input model.SignupInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := model.User{
		Name:        input.Name,
		Password:    input.Password,
		Email:       input.Email,
		Phonenumber: input.Phonenumber,
		Country:     input.Country,
		Gender:      input.Gender,
	}

	_, err := user.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Success"})
}

func Login(context *gin.Context) {
	var input model.LoginInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := model.FindUserByName(input.Name)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.ValidatePassword(input.Password)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwt, err := helper.GenerateJWT(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"access_token": jwt})
}
