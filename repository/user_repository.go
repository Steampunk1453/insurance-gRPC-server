package repository

import (
	"github.com/insurance-policy/insurancepolicy"
)

type UserRepository interface {
	Save(users []*insurancepolicy.User) error
	GetByID(id int64) (*insurancepolicy.User, error)
	GetByMobileNumber(mobileNumber string) (*insurancepolicy.User, error)
}
