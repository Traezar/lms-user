package model

import (
	"html"
	"lms-user/database"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index"`
	Name      string    `gorm:"size:255;not null;unique" json:"name"`
	Email     string    `gorm:"size:255;not null;unique;" json:"email"`
	Password  string    `gorm:"size:255;not null;" json:"-"`
	ManagerID uint      `json:"manager_id"`
}

func (user *User) Save() (*User, error) {
	err := database.Database.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) Update() (*User, error) {
	err := database.Database.Where("id = ?", user.ID).Updates(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.Name = html.EscapeString(strings.TrimSpace(user.Name))
	return nil
}

func (user *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func FindUserByName(name string) (User, error) {
	var user User
	err := database.Database.Where("name=?", name).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func GetUserById(id uint) (User, error) {
	var user User
	err := database.Database.Where("id=?", id).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}
