package controller

import (
	"net/http"

	"github.com/FJericho/test-mnc/internal/helper"
	"github.com/FJericho/test-mnc/internal/model"
	"github.com/FJericho/test-mnc/internal/service"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	Register(context *gin.Context)
	Login(context *gin.Context)
}

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func (c *userController) Register(context *gin.Context) {
	var user model.User

	if err := context.ShouldBindJSON(&user); err != nil {
		errorHandler := helper.UnprocessibleEntity("Invalid JSON body")
		context.AbortWithStatusJSON(errorHandler.Status(), errorHandler)
		return
	}


	response, err := c.userService.Register(&user)

	if err != nil {
		context.AbortWithStatusJSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"id": response.ID,
		"username": response.Username,
		"balance": response.Balance,
	})
}

func (c *userController) Login(context *gin.Context){
	var userLogin model.Login

	if err := context.ShouldBindJSON(&userLogin); err != nil {
		errorHandler := helper.UnprocessibleEntity("Invalid JSON body")
		context.AbortWithStatusJSON(errorHandler.Status(), errorHandler)
		return
	}

	token, err := c.userService.Login(&userLogin)

	if err != nil {
		context.AbortWithStatusJSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": token})
}

func Logout(context *gin.Context){}