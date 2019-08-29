package models

// UserModel - User model struct
type UserModel struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	IsActive string `json:"is_active"`
}
