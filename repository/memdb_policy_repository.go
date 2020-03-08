package repository

import (
	"github.com/hashicorp/go-memdb"
	"github.com/insurance-policy/insurancepolicy"
	"github.com/insurance-policy/utils"
)

type MemdbPolicyRepository struct {
	DB *memdb.MemDB
}

func NewMemdbPolicyRepository() PolicyRepository {
	return &MemdbPolicyRepository{
		DB: db,
	}
}

func (m MemdbPolicyRepository) Save(policies []*insurancepolicy.Policy) error {
	txn := GetDB().Txn(true)
	for _, p := range policies {
		p.MobileNumber = utils.GetNumberFormat(p.MobileNumber)
		err := txn.Insert("policy", p)
		if err != nil {
			return err
		}
	}
	txn.Commit()
	return nil
}

func (m MemdbPolicyRepository) GetByMobileNumber(mobileNumber string) ([]*insurancepolicy.Policy, error) {
	policies := make([]*insurancepolicy.Policy, 0)
	txn := GetDB().Txn(false)
	defer txn.Abort()

	it, err := txn.Get("policy", "mobileNumber", mobileNumber)
	if err != nil {
		return nil, err
	}

	for obj := it.Next(); obj != nil; obj = it.Next() {
		p := obj.(*insurancepolicy.Policy)
		policies = append(policies, p)
	}
	return policies, nil
}

