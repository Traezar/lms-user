package model

type SignupInput struct {
	Name        string `form:"name" json:"name" binding:"required"`
	Email       string `form:"email" json:"email" binding:"required"`
	Password    string `form:"password" json:"password" binding:"required"`
	Country     string `form:"country" json:"country" binding:"required"`
	Phonenumber string `form:"phonenumber" json:"phonenumber" binding:"required"`
	Gender      string `form:"gender" json:"gender" binding:"required"`
}

type LoginInput struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
