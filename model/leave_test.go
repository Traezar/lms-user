package model

import (
	"lms-user/database"
	"log"
	"reflect"
	"testing"
	"time"

	"gorm.io/gorm"
)

func beforeTest() {
	err := database.TestDatabase.Create(&User{ID: 54, Email: "test", Name: "raj", Password: "raj", ManagerID: 4}).Error
	if err != nil {
		log.Fatalf("%s", err)
	}
	err = database.TestDatabase.Create(&User{ID: 76, Email: "test", Name: "raj2", Password: "raj", ManagerID: 4}).Error
	if err != nil {
		log.Fatalf("%s", err)
	}
}

func afterTest() {
	err := database.TestDatabase.Delete(&User{ID: 54, Email: "test", Name: "raj", Password: "raj", ManagerID: 0}).Error
	if err != nil {
		log.Fatalf("%s", err)
	}
	err = database.TestDatabase.Delete(&User{ID: 76, Email: "test", Name: "raj2", Password: "raj", ManagerID: 0}).Error
	if err != nil {
		log.Fatalf("%s", err)
	}
}

func TestLeave_Save(t *testing.T) {
	database.ConnectTest()
	database.TestDatabase.AutoMigrate(&User{}, &Leave{})
	beforeTest()

	type fields struct {
		ID            uint
		ApplicantID   uint
		ManagerID     uint
		StartDatetime time.Time
		EndDatetime   time.Time
		CreatedAt     time.Time
		UpdatedAt     time.Time
		Type          string
		Status        string
		DeletedAt     gorm.DeletedAt
	}
	tests := []struct {
		name    string
		fields  fields
		want    *Leave
		wantErr bool
	}{
		{
			name: "Valid leave creation",
			fields: fields{
				ApplicantID:   1,
				ManagerID:     4,
				StartDatetime: time.Now().Add(time.Hour * 24),
				EndDatetime:   time.Now().Add(time.Hour * 48),
				Type:          "Holiday",
				Status:        "PENDING",
			},
			want:    &Leave{}, // We don't care about the generated ID
			wantErr: false,
		},
		{
			name: "Missing applicant ID",
			fields: fields{
				ManagerID:     4,
				StartDatetime: time.Now().Add(time.Hour * 24),
				EndDatetime:   time.Now().Add(time.Hour * 48),
				Type:          "Holiday",
				Status:        "PENDING",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Invalid date range (start after end)",
			fields: fields{
				ApplicantID:   1,
				ManagerID:     4,
				StartDatetime: time.Now().Add(time.Hour * 48),
				EndDatetime:   time.Now().Add(time.Hour * 24),
				Type:          "Holiday",
				Status:        "PENDING",
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			leave := &Leave{
				ID:            tt.fields.ID,
				ApplicantID:   tt.fields.ApplicantID,
				ManagerID:     tt.fields.ManagerID,
				StartDatetime: tt.fields.StartDatetime,
				EndDatetime:   tt.fields.EndDatetime,
				CreatedAt:     tt.fields.CreatedAt,
				UpdatedAt:     tt.fields.UpdatedAt,
				Type:          tt.fields.Type,
				Status:        tt.fields.Status,
				DeletedAt:     tt.fields.DeletedAt,
			}
			got, err := leave.Save()
			if (err != nil) != tt.wantErr {
				t.Errorf("Leave.Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Leave.Save() = %v, want %v", got, tt.want)
			}
		})
	}
	afterTest()
}

func TestLeave_Update(t *testing.T) {
	type fields struct {
		ID            uint
		ApplicantID   uint
		ManagerID     uint
		StartDatetime time.Time
		EndDatetime   time.Time
		CreatedAt     time.Time
		UpdatedAt     time.Time
		Type          string
		Status        string
		DeletedAt     gorm.DeletedAt
	}
	tests := []struct {
		name    string
		fields  fields
		want    *Leave
		wantErr bool
	}{{

		name: "Update leave status",
		fields: fields{
			ID:            1, // Existing leave ID
			ApplicantID:   1,
			ManagerID:     2,
			StartDatetime: time.Now().Add(time.Hour * 24),
			EndDatetime:   time.Now().Add(time.Hour * 48),
			Type:          "Holiday",
			Status:        "APPROVED", // Update status
		},
		want:    &Leave{ID: 1, Status: "APPROVED"}, // Only check ID and updated status
		wantErr: false,
	},
		{
			name: "Update leave with invalid date range",
			fields: fields{
				ID:            1, // Existing leave ID
				ApplicantID:   1,
				ManagerID:     2,
				StartDatetime: time.Now().Add(time.Hour * 48),
				EndDatetime:   time.Now().Add(time.Hour * 24),
				Type:          "Holiday",
				Status:        "PENDING",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Update non-existent leave",
			fields: fields{
				ID:            999, // Non-existent ID
				ApplicantID:   1,
				ManagerID:     2,
				StartDatetime: time.Now().Add(time.Hour * 24),
				EndDatetime:   time.Now().Add(time.Hour * 48),
				Type:          "Holiday",
				Status:        "APPROVED",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			leave := &Leave{
				ID:            tt.fields.ID,
				ApplicantID:   tt.fields.ApplicantID,
				ManagerID:     tt.fields.ManagerID,
				StartDatetime: tt.fields.StartDatetime,
				EndDatetime:   tt.fields.EndDatetime,
				CreatedAt:     tt.fields.CreatedAt,
				UpdatedAt:     tt.fields.UpdatedAt,
				Type:          tt.fields.Type,
				Status:        tt.fields.Status,
				DeletedAt:     tt.fields.DeletedAt,
			}
			got, err := leave.Update()
			if (err != nil) != tt.wantErr {
				t.Errorf("Leave.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Leave.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}
