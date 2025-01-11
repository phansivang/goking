package routes

import (
	"github.com/gin-gonic/gin"
	"goking/controllers"
)

func SetupUserRoutes(router *gin.Engine, userController *controllers.UserController) {
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/", userController.CreateUser)
		userRoutes.GET("/:id", userController.GetUserByID)
		userRoutes.PUT("/:id", userController.UpdateUser)
		userRoutes.GET("/", userController.ListUsers)
	}
}
