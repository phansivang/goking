package dtos

type CreateUserProfileRequest struct {
	UserID uint   `json:"user_id" binding:"required"`
	Bio    string `json:"bio"`
}

type UpdateUserProfileRequest struct {
	Bio string `json:"bio"`
}

type UserProfileResponse struct {
	ID     uint   `json:"id"`
	UserID uint   `json:"user_id"`
	Bio    string `json:"bio"`
}
