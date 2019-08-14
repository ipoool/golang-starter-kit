package models

// UserTokenModel - Model for response token
type UserTokenModel struct {
	ID          int    `json:"id"`
	Email       string `json:"email"`
	DisplayName string `json:"display_name"`
	RoleName    string `json:"role"`
}
