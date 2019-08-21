package models

import "time"

// UserTokenModel - Model for response token
type UserTokenModel struct {
	ID          int    `json:"id"`
	Email       string `json:"email"`
	DisplayName string `json:"display_name"`
	RoleName    string `json:"role"`
}

// UserModel - Model for user
type UserModel struct {
	ID             int        `json:"id"`
	Name           string     `json:"name"`
	Email          string     `json:"email"`
	BirthPlace     string     `json:"birth_place"`
	BirthDate      string     `json:"birth_date"`
	RoleID         int        `json:"role_id"`
	IsActive       bool       `json:"is_active"`
	Phone          string     `json:"phone"`
	IsVerified     bool       `json:"is_verified"`
	VerifiedDate   string     `json:"verified_date"`
	ProfilePicture string     `json:"profile_picture"`
	CreatedAt      *time.Time `json:"created_at"`
}

// UserAuth - Model for user authentication
type UserAuth struct {
	UserID     int    `json:"user_id"`
	FacebookID string `json:"facebook_id"`
	TwitterID  string `json:"twitter_id"`
	GoogleID   string `json:"google_id"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}
