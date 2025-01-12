package models

type UserProfile struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	UserID uint   `gorm:"not null" json:"user_id"`
	Bio    string `gorm:"type:text" json:"bio"`
}
