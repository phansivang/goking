package dtos

import "time"

type CreateUserRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type UpdateUserRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UserResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type UserWithProfileResponse struct {
	ID          uint                `json:"id"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	CreatedAt   time.Time           `json:"created_at"`
	UserProfile UserProfileResponse `json:"user_profile"`
}

type ListUsersRequest struct {
	ID          int       `form:"id"`
	Name        string    `form:"name"`
	Description string    `form:"description"`
	StartDate   time.Time `form:"start_date" time_format:"2006-01-02"`
	EndDate     time.Time `form:"end_date" time_format:"2006-01-02"`
	Page        int       `form:"page,default=1"`
	Limit       int       `form:"limit,default=10"`
}
