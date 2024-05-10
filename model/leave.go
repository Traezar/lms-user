package model

import (
	"lms-user/database"
	"time"

	"gorm.io/gorm"
)

type Leave struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	ApplicantID   uint           `json:"applicant_id"`
	ManagerID     uint           `json:"manager_id"`
	StartDatetime time.Time      `gorm:"not null;" json:"start_datetime"`
	EndDatetime   time.Time      `gorm:"not null;" json:"end_datetime"`
	CreatedAt     time.Time      `gorm:"default:CURRENT_TIMESTAMP"  json:"-"`
	UpdatedAt     time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"-"`
	Type          string         `gorm:"size:255;not null;" json:"type"`
	Status        string         `gorm:"size:255;not null;" json:"status"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

func (leave *Leave) Save() (*Leave, error) {
	err := database.Database.Create(&leave).Error
	if err != nil {
		return &Leave{}, err
	}
	return leave, nil
}

func (leave *Leave) Update() (*Leave, error) {
	err := database.Database.UpdateColumns(&leave).Error
	if err != nil {
		return &Leave{}, err
	}
	return leave, nil
}
