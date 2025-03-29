package router

import (
	"log"
	"os"

	"github.com/FJericho/test-mnc/internal/controller"
	"github.com/FJericho/test-mnc/internal/database"
	"github.com/FJericho/test-mnc/internal/middleware"
	"github.com/FJericho/test-mnc/internal/repository"
	"github.com/FJericho/test-mnc/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
	err error
)

func StartServer() {
	db, err = database.ConnectDB()

	if err != nil {
		log.Fatal(err)
	}
	router := gin.Default()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	userRouter := router.Group("/users")
	{
		userRouter.Use(middleware.Authentication())
		userRouter.POST("/register", userController.Register)
		userRouter.POST("/login", userController.Login)
	}
	var PORT = os.Getenv("APP_PORT")
	router.Run(":" + PORT)
}