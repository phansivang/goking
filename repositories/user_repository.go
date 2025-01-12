package repositories

import (
	"goking/dtos"
	"goking/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.Preload("UserProfile").First(&user, id).Error
	return &user, err
}

func (r *UserRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}

func (r *UserRepository) FindAll(req dtos.ListUsersRequest) []models.User {
	var users []models.User
	offset := (req.Page - 1) * req.Limit
	query := r.db.Model(&models.User{})

	/* FILTER CRITERIA */
	if req.ID != 0 {
		query = query.Where("id = ?", req.ID)
	}

	if req.Name != "" {
		query = query.Where("name LIKE ?", "%"+req.Name+"%")
	}

	if req.Description != "" {
		query = query.Where("description LIKE ?", "%"+req.Description+"%")
	}
	if !req.StartDate.IsZero() && !req.EndDate.IsZero() {
		query = query.Where("created_at BETWEEN ? AND ?", req.StartDate, req.EndDate)
	}

	/* PAGINATION */
	query.Offset(offset)
	query.Limit(req.Limit)
	query.Find(&users)

	return users
}
