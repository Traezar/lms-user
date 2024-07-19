package model

import (
	"reflect"
	"testing"
)

func TestGetUserById(t *testing.T) {
	type args struct {
		id uint
	}
	tests := []struct {
		name    string
		args    args
		want    User
		wantErr bool
	}{
		{
			name: "Find existing user by ID",
			args: args{id: 1},
			want: User{ID: 1},
		},
		{
			name:    "User not found",
			args:    args{id: 999}, // Non-existent ID
			want:    User{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUserById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindUserByName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    User
		wantErr bool
	}{
		{
			name: "Find existing user by name",
			args: args{name: "John Doe"},
			want: User{Name: "John Doe"},
		},
		{
			name:    "User not found",
			args:    args{name: "Nonexistent User"},
			want:    User{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindUserByName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindUserByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindUserByName() = %v, want %v", got, tt.want)
			}
		})
	}
}
