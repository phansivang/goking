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
	ID          int       `form:"id"`                                  // Filter by id (partial match)
	Name        string    `form:"name"`                                // Filter by name (partial match)
	Description string    `form:"description"`                         // Filter by description (partial match)
	StartDate   time.Time `form:"start_date" time_format:"2006-01-02"` // Filter by created_at (start date)
	EndDate     time.Time `form:"end_date" time_format:"2006-01-02"`   // Filter by created_at (end date)
	Page        int       `form:"page,default=1"`                      // Pagination: page number
	Limit       int       `form:"limit,default=10"`                    // Pagination: number of items per page
}
