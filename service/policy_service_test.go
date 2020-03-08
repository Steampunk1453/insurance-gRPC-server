package service

import (
	"github.com/insurance-policy/insurancepolicy"
	"github.com/insurance-policy/repository"
	"reflect"
	"testing"
)

func TestNewPolicyService(t *testing.T) {
	type args struct {
		policyRepository repository.PolicyRepository
	}
	var tests []struct {
		name string
		args args
		want *policyService
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPolicyService(tt.args.policyRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPolicyService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_policyService_GetByMobileNumber(t *testing.T) {
	type fields struct {
		policyRepository repository.PolicyRepository
	}
	type args struct {
		mobileNumber string
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    []*insurancepolicy.Policy
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := policyService{
				policyRepository: tt.fields.policyRepository,
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

func Test_policyService_Save(t *testing.T) {
	type fields struct {
		policyRepository repository.PolicyRepository
	}
	type args struct {
		users []*insurancepolicy.Policy
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := policyService{
				policyRepository: tt.fields.policyRepository,
			}
			if err := s.Save(tt.args.users); (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}