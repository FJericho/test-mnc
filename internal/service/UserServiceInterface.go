package service

import (
	"github.com/FJericho/test-mnc/internal/helper"
	"github.com/FJericho/test-mnc/internal/model"
)



type UserService interface {
	Register(*model.User) (*model.User, helper.Error)
	Login(account *model.Login) (string, helper.Error)
}