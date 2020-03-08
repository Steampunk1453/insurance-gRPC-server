package service

import (
	"github.com/insurance-policy/insurancepolicy"
	"github.com/insurance-policy/repository"
	"reflect"
	"testing"
)

func TestNewUserService(t *testing.T) {
	type args struct {
		userRepository repository.UserRepository
	}
	var tests []struct {
		name string
		args args
		want *userService
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserService(tt.args.userRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_GetByID(t *testing.T) {
	type fields struct {
		userRepository repository.UserRepository
	}
	type args struct {
		id int64
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    *insurancepolicy.User
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := userService{
				userRepository: tt.fields.userRepository,
			}
			got, err := s.GetByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_GetByMobileNumber(t *testing.T) {
	type fields struct {
		userRepository repository.UserRepository
	}
	type args struct {
		mobileNumber string
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    *insurancepolicy.User
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := userService{
				userRepository: tt.fields.userRepository,
			}
			got, err := s.GetByMobileNumber(tt.args.mobileNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByMobileNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByMobileNumber() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_Save(t *testing.T) {
	type fields struct {
		userRepository repository.UserRepository
	}
	type args struct {
		users []*insurancepolicy.User
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := userService{
				userRepository: tt.fields.userRepository,
			}
			if err := s.Save(tt.args.users); (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}