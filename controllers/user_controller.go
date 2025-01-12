package controllers

import (
	"github.com/gin-gonic/gin"
	"goking/dtos"
	"goking/services"
	"goking/utils"
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
	utils.BindJson(ctx, &createUserReq)
	c.userService.CreateUser(ctx, createUserReq)
}

func (c *UserController) GetUserByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	c.userService.GetUserByID(id, ctx)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var updateUserReq dtos.UpdateUserRequest
	utils.BindJson(ctx, &updateUserReq)

	c.userService.UpdateUser(ctx, uint(id), updateUserReq)
}

func (c *UserController) ListUsers(ctx *gin.Context) {
	var listUsersReq dtos.ListUsersRequest
	utils.BindQuery(ctx, &listUsersReq)
	c.userService.ListUsers(ctx, listUsersReq)
}
