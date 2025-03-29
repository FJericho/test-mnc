package repository

import (
	"github.com/FJericho/test-mnc/internal/helper"
	"github.com/FJericho/test-mnc/internal/model"
	"gorm.io/gorm"
)

type userRepositoryImpl struct{
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

func(r *userRepositoryImpl) Register(user *model.User)(*model.User, helper.Error){
	err := r.db.Create(&user).Error

	if err != nil {
		return nil, helper.InternalServerError("Something went wrong")
	}

	return user, nil

}

func(r *userRepositoryImpl) Login(userLogin *model.Login)(*model.User, helper.Error){
	var user model.User
	
	err := r.db.Where("username = ?", userLogin.Username).First(&user).Error

	if err != nil {
		return nil, helper.Unauthorized("Invalid email/password")
	}

	return &user, nil
}