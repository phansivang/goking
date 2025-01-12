package routes

import (
	"github.com/gin-gonic/gin"
	"goking/pkg/controllers"
)

func SetupRoutes(router *gin.Engine, userController *controllers.UserController) {
	SetupUserRoutes(router, userController)
	SetupRuntimeRoutes(router)

	// Add more route groups here for other controllers (e.g., profile, auth, etc.)
}
