package service

import (
	"github.com/insurance-policy/insurancepolicy"
	"github.com/insurance-policy/repository"
)

type PolicyService interface {
	Save(users []*insurancepolicy.Policy) error
	GetByMobileNumber(mobileNumber string) ([]*insurancepolicy.Policy, error)
}

type policyService struct {
	policyRepository repository.PolicyRepository
}

func NewPolicyService(policyRepository repository.PolicyRepository) *policyService {
	return &policyService{
		policyRepository: policyRepository,
	}
}

func (s policyService) Save(users []*insurancepolicy.Policy) error {
	err := s.policyRepository.Save(users)
	if err != nil {
		return err
	}
	return nil
}

func (s policyService) GetByMobileNumber(mobileNumber string) ([]*insurancepolicy.Policy, error) {
	policies, err := s.policyRepository.GetByMobileNumber(mobileNumber)
	if err != nil {
		return nil, err
	}
	return policies, nil
}

