package services

import (
	"github.com/gin-gonic/gin"
	"goking/pkg/dtos"
	"goking/pkg/models"
	"goking/pkg/repositories"
	utils2 "goking/pkg/utils"
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
	utils2.SuccessResponse(ctx)
}

func (s *UserService) GetUserByID(id int, ctx *gin.Context) {
	user, err := s.userRepo.FindByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils2.ErrorResponse{Message: "user not found", StatusCode: http.StatusNotFound})
		return
	}

	ctx.JSON(http.StatusOK, utils2.Responses{Data: user, Message: "OK"})
}

func (s *UserService) UpdateUser(ctx *gin.Context, id uint, req dtos.UpdateUserRequest) {
	user, _ := s.userRepo.FindByID(id)

	err := s.userRepo.Update(user)
	if err != nil {
		return
	}
	utils2.SuccessResponse(ctx)
}

func (s *UserService) ListUsers(ctx *gin.Context, req dtos.ListUsersRequest) {
	limit, offset := req.Limit, req.Page

	users := s.userRepo.FindAll(req)

	ctx.JSON(http.StatusOK, utils2.Responses{Data: users, Message: "SUCCESS", Limit: limit, Offset: offset})
}
