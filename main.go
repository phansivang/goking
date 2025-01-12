package main

import (
	"github.com/gin-gonic/gin"
	"goking/pkg/config"
	"goking/pkg/controllers"
	models2 "goking/pkg/models"
	"goking/pkg/repositories"
	"goking/pkg/routes"
	"goking/pkg/services"
)

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&models2.User{}, &models2.UserProfile{})

	userRepo := repositories.NewUserRepository(config.DB)

	userService := services.NewUserService(userRepo)

	userController := controllers.NewUserController(userService)

	router := gin.Default()

	//router.Use(middleware.ResponseInterceptor())

	routes.SetupRoutes(router, userController)

	router.Run(":8080")
}
