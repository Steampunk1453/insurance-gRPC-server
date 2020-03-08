package repository

import (
	"github.com/hashicorp/go-memdb"
	"github.com/insurance-policy/insurancepolicy"
	"github.com/insurance-policy/utils"
)

type memdbUserRepository struct {
	DB *memdb.MemDB
}

func NewMemdbUserRepository() *memdbUserRepository {
	return &memdbUserRepository{
		DB: db,
	}
}

func (m memdbUserRepository) Save(users []*insurancepolicy.User) error {
	txn := GetDB().Txn(true)
	for _, u := range users {
		u.MobileNumber = utils.GetNumberFormat(u.MobileNumber)
		err := txn.Insert("user", u)
		if err != nil {
			return err
		}
	}
	txn.Commit()
	return nil
}

func (m memdbUserRepository) GetByID(id int64) (*insurancepolicy.User, error) {
	txn := GetDB().Txn(false)
	defer txn.Abort()
	raw, err := txn.First("user", "id", id)
	if err != nil {
		return nil, err
	}
	return raw.(*insurancepolicy.User), nil
}

func (m memdbUserRepository) GetByMobileNumber(mobileNumber string) (*insurancepolicy.User, error) {
	txn := GetDB().Txn(false)
	defer txn.Abort()
	raw, err := txn.First("user", "mobileNumber", mobileNumber)
	if err != nil {
		return nil, err
	}
	return raw.(*insurancepolicy.User), nil
}




