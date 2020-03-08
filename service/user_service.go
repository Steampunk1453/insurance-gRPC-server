package service

import (
	"github.com/insurance-policy/insurancepolicy"
	"github.com/insurance-policy/repository"
)

type UserService interface {
	Save(users []*insurancepolicy.User) error
	GetByID(id int64) (*insurancepolicy.User, error)
	GetByMobileNumber(mobileNumber string) (*insurancepolicy.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *userService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s userService) Save(users []*insurancepolicy.User) error {
	err := s.userRepository.Save(users)
	if err != nil {
		return err
	}
	return nil
}

func (s userService) GetByID(id int64) (*insurancepolicy.User, error) {
	user, err := s.userRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s userService) GetByMobileNumber(mobileNumber string) (*insurancepolicy.User, error) {
	user, err := s.userRepository.GetByMobileNumber(mobileNumber)
	if err != nil {
		return nil, err
	}
	return user, nil
}

