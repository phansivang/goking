package services

import (
	"github.com/gin-gonic/gin"
	"goking/dtos"
	"goking/models"
	"goking/repositories"
	"goking/utils"
	"net/http"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(ctx *gin.Context, req dtos.CreateUserRequest) {
	user := models.User{
		Name:        req.Name,
		Description: req.Description,
	}

	err := s.userRepo.Create(&user)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessResponses{Message: "SUCCESS", StatusCode: http.StatusOK})
}

func (s *UserService) GetUserByID(id int, ctx *gin.Context) {
	user, err := s.userRepo.FindByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse{Message: "user not found", StatusCode: http.StatusNotFound})
		return
	}

	ctx.JSON(http.StatusOK, utils.Responses{Data: user, Message: "OK"})
}

func (s *UserService) UpdateUser(id uint, req dtos.UpdateUserRequest) (*models.User, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Description != "" {
		user.Description = req.Description
	}

	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) ListUsers(ctx *gin.Context, req dtos.ListUsersRequest) {
	limit, offset := req.Limit, req.Page

	users := s.userRepo.FindAll(req)

	ctx.JSON(http.StatusOK, utils.Responses{Data: users, Message: "SUCCESS", Limit: limit, Offset: offset})
}
