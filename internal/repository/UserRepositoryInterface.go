package repository

import (
	"github.com/FJericho/test-mnc/internal/helper"
	"github.com/FJericho/test-mnc/internal/model"
)

type UserRepository interface{
	 	Register(user *model.User)(*model.User, helper.Error)
		Login(userLogin *model.Login)(*model.User, helper.Error)
}