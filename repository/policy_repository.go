package repository

import (
	"github.com/insurance-policy/insurancepolicy"
)

type PolicyRepository interface {
	Save(users []*insurancepolicy.Policy) error
	GetByMobileNumber(mobileNumber string) ([]*insurancepolicy.Policy, error)
}

