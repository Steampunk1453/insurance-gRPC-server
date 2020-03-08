package service

import (
	"github.com/insurance-policy/insurancepolicy"
	"reflect"
	"testing"
)

func TestNewAmsService(t *testing.T) {
	var tests []struct {
		name string
		want *amsService
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAmsService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAmsService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_amsService_GetPoliciesByAgentID(t *testing.T) {
	type args struct {
		id int64
	}
	var tests []struct {
		name string
		args args
		want []*insurancepolicy.Policy
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &amsService{}
			if got := a.GetPoliciesByAgentID(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPoliciesByAgentID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_amsService_GetUsersByAgentID(t *testing.T) {
	type args struct {
		id int64
	}
	var tests []struct {
		name string
		args args
		want []*insurancepolicy.User
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &amsService{}
			if got := a.GetUsersByAgentID(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUsersByAgentID() = %v, want %v", got, tt.want)
			}
		})
	}
}