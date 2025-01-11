package main

import (
	"github.com/gin-gonic/gin"
	"goking/config"
	"goking/controllers"
	"goking/models"
	"goking/repositories"
	"goking/routes"
	"goking/services"
)

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{}, &models.UserProfile{})

	userRepo := repositories.NewUserRepository(config.DB)

	userService := services.NewUserService(userRepo)

	userController := controllers.NewUserController(userService)

	router := gin.Default()

	//router.Use(middleware.ResponseInterceptor())

	routes.SetupRoutes(router, userController)

	router.Run(":8080")
}
