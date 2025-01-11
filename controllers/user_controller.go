package controllers

import (
	"github.com/gin-gonic/gin"
	"goking/dtos"
	"goking/services"
	"net/http"
	"strconv"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var createUserReq dtos.CreateUserRequest
	if err := ctx.ShouldBindJSON(&createUserReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.userService.CreateUser(createUserReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userResponse := dtos.UserResponse{
		Name:        user.Name,
		Description: user.Description,
		CreatedAt:   user.CreatedAt,
	}

	ctx.Set("Data", userResponse)
}

func (c *UserController) GetUserByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	c.userService.GetUserByID(id, ctx)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var updateUserReq dtos.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&updateUserReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.userService.UpdateUser(uint(id), updateUserReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userResponse := dtos.UserResponse{
		ID:          user.ID,
		Name:        user.Name,
		Description: user.Description,
		CreatedAt:   user.CreatedAt,
	}

	ctx.JSON(http.StatusOK, userResponse)
}

func (c *UserController) ListUsers(ctx *gin.Context) {
	var listUsersReq dtos.ListUsersRequest
	c.userService.ListUsers(ctx, listUsersReq)
}
