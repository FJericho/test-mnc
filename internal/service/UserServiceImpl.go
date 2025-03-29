package service

import (
	"github.com/FJericho/test-mnc/internal/helper"
	"github.com/FJericho/test-mnc/internal/model"
	"github.com/FJericho/test-mnc/internal/repository"
	"github.com/asaskevich/govalidator"
)

type userServiceImpl struct{
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService{
	return &userServiceImpl{
		userRepo: userRepo,
	}
}

func (s *userServiceImpl)Register(user *model.User) (*model.User, helper.Error){
	if _, err := govalidator.ValidateStruct(user); err != nil {
		return nil, helper.BadRequest(err.Error())
	}

	password, err := helper.HashPassword(user.Password)

	if err != nil {
		return nil, err
	}

	user.Password = password
	response, err := s.userRepo.Register(user)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *userServiceImpl)Login(account *model.Login) (string, helper.Error){
	if _, err := govalidator.ValidateStruct(account); err != nil {
		return "", helper.BadRequest(err.Error())
	}

	response, err := s.userRepo.Login(account)

	if err != nil {
		return "", err
	}

	if isPasswordCorrect := helper.ComparePassword(account.Password, response.Password); !isPasswordCorrect {
		return "", helper.Unauthorized("Invalid username/password")
	}

	token, err := helper.GenerateToken(response.ID, response.Fullname)

	if err != nil {
		return "", err
	}

	return token, nil
}

