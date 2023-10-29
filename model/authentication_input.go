package model

type AuthenticationInput struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Country     string `json:"country" binding:"required"`
	Phonenumber string `json:"phonenumber" binding:"required"`
	Gender      string `json:"gender" binding:"required"`
}
