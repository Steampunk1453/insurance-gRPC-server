package service

import (
	"encoding/json"
	"github.com/insurance-policy/insurancepolicy"
	"github.com/insurance-policy/restclient"
	"log"
)

type AmsService interface {
	GetUsersByAgentID(id int64) []*insurancepolicy.User
	GetPoliciesByAgentID(id int64) []*insurancepolicy.Policy
}

type amsService struct{
}

func NewAmsService() *amsService {
	return &amsService{
	}
}

func (a *amsService) GetUsersByAgentID(id int64) []*insurancepolicy.User {
		var data []byte
		users := make([]*insurancepolicy.User, 0)
		data = restclient.GetUsers(id)
		if err := json.Unmarshal(data, &users); err != nil {
			log.Fatalf("Failed to get default users: %v", err)
		}
	return users
}

func (a *amsService) GetPoliciesByAgentID(id int64) []*insurancepolicy.Policy {
	var data []byte
	policies := make([]*insurancepolicy.Policy, 0)
	data = restclient.GetPolicies(id)
	if err := json.Unmarshal(data, &policies); err != nil {
		log.Fatalf("Failed to get default policies: %v", err)
	}
	return policies
}
